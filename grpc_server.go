package rest4grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tmarcus87/rest4grpc/logger"
	"github.com/tmarcus87/rest4grpc/pb"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type GrpcServer struct {
	native *grpc.Server
}

func (s *GrpcServer) Start(port uint16) error {
	logger.Info("Starting server")

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		return err
	}

	serverOpts := make([]grpc.ServerOption, 0)
	serverOpts = append(serverOpts, grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	s.native = grpc.NewServer(serverOpts...)
	pb.RegisterTestServiceServer(s.native, NewTestServiceServer())
	reflection.Register(s.native)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return s.serveAndWait(
		func() error {
			return s.native.Serve(listener)
		},
		func() bool {
			cc, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
			if err != nil {
				return false
			}

			c := pb.NewTestServiceClient(cc)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			if _, err := c.Live(ctx, &empty.Empty{}); err != nil {
				return false
			}

			return true
		},
		5*time.Second,
		100*time.Millisecond)
}

func (s *GrpcServer) serveAndWait(
	serveFn func() error,
	watchFn func() bool,
	watchTimeout time.Duration,
	watchInterval time.Duration) error {

	var (
		errCh = make(chan error)
		ready = errors.New("ok")
	)

	go func() {
		errCh <- serveFn()
	}()

	go func() {
		for {
			if watchFn() {
				errCh <- ready
			}
			time.Sleep(watchInterval)
		}
	}()

	for i := 0; i < int(watchTimeout/watchInterval); i++ {
		select {
		case err := <-errCh:
			if err == ready {
				return nil
			}
			return err
		case <-time.After(watchInterval):
			// Do nothing
		}
	}
	return errors.New("timeout")
}

func (s *GrpcServer) Stop() {
	logger.Info("Stopping server.")
	s.native.GracefulStop()
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

type TestServiceServer struct {
}

func (s TestServiceServer) Live(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s TestServiceServer) MethodA(ctx context.Context, in *pb.ParamA) (*pb.ResponseA, error) {
	traceID := trace.FromContext(ctx).SpanContext().TraceID.String()
	spanID := trace.FromContext(ctx).SpanContext().SpanID.String()

	return &pb.ResponseA{
		TraceId:      traceID,
		ParentSpanId: spanID,
		Name:         in.Name,
		Age:          in.Age,
	}, nil
}

func (s TestServiceServer) MethodB(ctx context.Context, in *pb.ParamB) (*pb.ResponseB, error) {
	traceID := trace.FromContext(ctx).SpanContext().TraceID.String()

	return &pb.ResponseB{
		TraceId: traceID,
		Name:    in.Name,
		Address: in.Address,
	}, nil
}

func NewTestServiceServer() pb.TestServiceServer {
	return &TestServiceServer{}
}
