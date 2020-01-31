package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/tmarcus87/rest4grpc/logger"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	reflection "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
	"io"
	"net/url"
	"strings"
)

const tracingHeader = "grpc-trace-bin"

var (
	MethodDescriptorFetchError = errors.New("failed to fetch method descriptor")
	MethodDescriptorNotFound   = errors.New("method descriptor is not found")
)

type DynamicGrpcClient struct {
	cc   *grpc.ClientConn
	refc *grpcreflect.Client

	refCtx    context.Context
	refCancel context.CancelFunc
}

func NewDynamicGrpcClient(target string) (*DynamicGrpcClient, error) {
	client := DynamicGrpcClient{}

	if !strings.Contains(target, "://") {
		return nil, fmt.Errorf("invalid target format")
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target url : %w", err)
	}

	opts := make([]grpc.DialOption, 0)

	if targetURL.Scheme == "grpc" {
		opts = append(opts, grpc.WithInsecure())
	} else {
		return nil, fmt.Errorf("unsupported scheme(%s)", targetURL.Scheme)
	}

	// Connect to server
	client.cc, err = grpc.Dial(targetURL.Host, opts...)
	if err != nil {
		return nil, err
	}

	// Create context/cancelFunc
	client.refCtx, client.refCancel = context.WithCancel(context.Background())

	// Create reflection client
	client.refc = grpcreflect.NewClient(client.refCtx, reflection.NewServerReflectionClient(client.cc))

	return &client, nil
}

func (c *DynamicGrpcClient) Invoke(ctx context.Context, service, method string, req io.Reader) ([]byte, error) {
	mtd, err := c.getMethodDesc(ctx, service, method)
	if err != nil {
		return nil, err
	}

	mf := dynamic.NewMessageFactoryWithDefaults()
	reqMsg := mf.NewMessage(mtd.GetInputType())
	if err := jsonpb.Unmarshal(req, reqMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal : %w", err)
	}

	stub := grpcdynamic.NewStubWithMessageFactory(c.cc, mf)

	md := make(metadata.MD)
	md[tracingHeader] =
		append(
			md[tracingHeader],
			string(propagation.Binary(trace.FromContext(ctx).SpanContext())))

	invokeCtx := metadata.NewOutgoingContext(ctx, md)

	resMsg, err := stub.InvokeRpc(invokeCtx, mtd, reqMsg)
	if err != nil {
		return nil, err
	}

	return json.Marshal(resMsg)
}

func (c *DynamicGrpcClient) getMethodDesc(ctx context.Context, service, method string) (*desc.MethodDescriptor, error) {
	logger.FromContext(ctx).
		Debug("Lookup method descriptor",
			zap.String("service", service),
			zap.String("method", method))

	sd, err := c.refc.ResolveService(service)
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if ok {
			logger.FromContext(ctx).
				Warn("failed to fetch method descriptor",
					zap.String("code", grpcStatus.Code().String()),
					zap.String("message", grpcStatus.Message()))
			return nil, MethodDescriptorFetchError
		}
		return nil, err
	}
	for _, m := range sd.GetMethods() {
		if m.GetName() == method {
			return m, nil
		}
	}
	return nil, MethodDescriptorNotFound
}

func (c *DynamicGrpcClient) Close() {
	c.refCancel()
}
