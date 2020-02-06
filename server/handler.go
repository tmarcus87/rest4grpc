package server

import (
	"github.com/joncalhoun/qson"
	"github.com/tmarcus87/rest4grpc/grpc"
	"github.com/tmarcus87/rest4grpc/logger"
	"github.com/tmarcus87/rest4grpc/message"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

//var (
//	code2status = make(map[codes.Code]int)
//)
//
//func init() {
//	code2status[codes.OK] = http.StatusOK
//	//code2status[codes.Canceled] = http.Status
//	code2status[codes.Unknown] = http.StatusInternalServerError
//	code2status[codes.InvalidArgument] = http.StatusBadRequest
//	//code2status[codes.DeadlineExceeded] = http.Status
//	code2status[codes.NotFound] = http.StatusNotFound
//	code2status[codes.AlreadyExists] = http.StatusConflict
//	code2status[codes.PermissionDenied] = http.StatusForbidden
//	code2status[codes.ResourceExhausted] = http.StatusForbidden
//	code2status[codes.FailedPrecondition] = http.StatusPreconditionFailed
//	code2status[codes.Aborted] = http.StatusInternalServerError
//	//code2status[codes.OutOfRange] = http.Status
//	code2status[codes.Unimplemented] = http.StatusForbidden
//	code2status[codes.Internal] = http.StatusInternalServerError
//	code2status[codes.Unavailable] = http.StatusServiceUnavailable
//	//code2status[codes.DataLoss] = http.StatusInternalServerError
//	code2status[codes.Unauthenticated] = http.StatusUnauthorized
//}

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
		}
		ctx.Response(http.StatusInternalServerError, err)
		return
	}
	ctx.Send(http.StatusOK, bytes)
}
