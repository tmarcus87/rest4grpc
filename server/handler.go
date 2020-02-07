package server

import (
	"github.com/joncalhoun/qson"
	"github.com/tmarcus87/rest4grpc/grpc"
	"github.com/tmarcus87/rest4grpc/logger"
	"github.com/tmarcus87/rest4grpc/message"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

var (
	code2status = map[codes.Code]int{
		codes.OK:                 http.StatusOK,
		codes.Canceled:           499, // ClientClosedRequest
		codes.Unknown:            http.StatusInternalServerError,
		codes.InvalidArgument:    http.StatusBadRequest,
		codes.DeadlineExceeded:   http.StatusGatewayTimeout,
		codes.NotFound:           http.StatusNotFound,
		codes.AlreadyExists:      http.StatusConflict,
		codes.PermissionDenied:   http.StatusForbidden,
		codes.ResourceExhausted:  http.StatusTooManyRequests,
		codes.FailedPrecondition: http.StatusBadRequest,
		codes.Aborted:            http.StatusConflict,
		codes.OutOfRange:         http.StatusBadRequest,
		codes.Unimplemented:      http.StatusNotImplemented,
		codes.Internal:           http.StatusInternalServerError,
		codes.Unavailable:        http.StatusServiceUnavailable,
		codes.DataLoss:           http.StatusInternalServerError,
		codes.Unauthenticated:    http.StatusUnauthorized,
	}
)

type ProxyHandler struct {
	client           *grpc.DynamicGrpcClient
	samplingFraction float64
}

func NewProxyHandler(client *grpc.DynamicGrpcClient, samplingFraction float64) (*ProxyHandler, error) {
	return &ProxyHandler{
		client:           client,
		samplingFraction: samplingFraction,
	}, nil
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewRequestContext(r, w, h.samplingFraction)

	logger.FromContext(ctx).
		Debug("RequestInfo",
			zap.String("RequestURI", r.RequestURI),
			zap.String("URL.Path", r.URL.Path))

	// Convert path to service&method
	service, method := path.Split(r.URL.Path)
	service = strings.TrimSuffix(strings.TrimPrefix(service, "/"), "/")
	if service == "" || method == "" {
		ctx.Send(http.StatusNotFound, "Not found")
		return
	}
	service = service[1:]

	var msg message.Message
	switch m := r.Method; m {
	case http.MethodPost:
		bytes, err := ioutil.ReadAll(ctx.request.Body)
		defer ctx.request.Body.Close()
		if err != nil {
			logger.FromContext(ctx).Warn("Failed to read client body", zap.Error(err))
			ctx.Response(http.StatusInternalServerError, err.Error())
			return
		}
		msg = message.NewJsonMessage(bytes)
		break

	case http.MethodGet:
		bytes, err := qson.ToJSON(ctx.request.URL.RawQuery)
		if err != nil {
			logger.FromContext(ctx).Debug("Failed to parse query", zap.Error(err))
			ctx.Response(http.StatusBadRequest, "Failed to parse query")
			return
		}
		msg = message.NewJsonMessage(bytes)
		break

	default:
		ctx.Send(http.StatusUnsupportedMediaType, "Method("+m+") is supported")
		return
	}

	// Invoke gRPC
	bytes, err := h.client.Invoke(ctx, service, method, msg)
	if err != nil {
		if err == grpc.MethodDescriptorNotFound {
			ctx.Send(http.StatusNotFound, "Method not found")
			return
		}

		// Response status conversion of gRPC to http
		hs := http.StatusInternalServerError
		if gs, ok := status.FromError(err); ok {
			hs = code2status[gs.Code()]
		}

		ctx.Response(hs, err)
		return
	}
	ctx.Send(http.StatusOK, bytes)
}
