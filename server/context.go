package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/tmarcus87/rest4grpc/logger"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"time"
)

const (
	b3TraceHeaderName   = "X-B3-TraceId"
	b3SpanHeaderName    = "X-B3-SpanId"
	b3SampledHeaderName = "X-B3-Sampled"
)

type RequestContext struct {
	ctx      context.Context
	request  *http.Request
	response http.ResponseWriter
}

func (c *RequestContext) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *RequestContext) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *RequestContext) Err() error {
	return c.ctx.Err()
}

func (c *RequestContext) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func NewRequestContext(req *http.Request, res http.ResponseWriter, samplingFraction float64) *RequestContext {
	ctx :=
		&RequestContext{
			ctx:      context.Background(),
			request:  req,
			response: res,
		}

	ctx.prepareTracing(samplingFraction)

	return ctx
}

func (c *RequestContext) prepareTracing(samplingFraction float64) {
	name := strings.TrimPrefix(c.request.URL.Path, "/")
	name = strings.Replace(name, "/", ".", -1)

	opts := make([]trace.StartOption, 0)
	opts = append(opts, trace.WithSpanKind(trace.SpanKindServer))
	opts = append(opts, trace.WithSampler(trace.ProbabilitySampler(samplingFraction)))

	if b3sc, ok := c.spanCtxFromB3(); ok {
		c.ctx, _ = trace.StartSpanWithRemoteParent(c.ctx, name, *b3sc, opts...)
		return
	}

	c.ctx, _ = trace.StartSpan(c.ctx, name)
}

func (c *RequestContext) spanCtxFromB3() (*trace.SpanContext, bool) {
	sc := trace.SpanContext{}

	b3TraceID := c.request.Header.Get(b3TraceHeaderName)
	b3SpanID := c.request.Header.Get(b3SpanHeaderName)
	b3Sampled := c.request.Header.Get(b3SampledHeaderName)

	if b3TraceID == "" || b3SpanID == "" {
		return nil, false
	}

	traceBytes, err := base64.StdEncoding.DecodeString(b3TraceID)
	if err != nil {
		return nil, false
	}
	copy(sc.TraceID[:], traceBytes)

	spanBytes, err := base64.StdEncoding.DecodeString(b3SpanID)
	if err != nil {
		return nil, false
	}
	copy(sc.SpanID[:], spanBytes)

	if b3Sampled == "1" {
		sc.TraceOptions = trace.TraceOptions(1)
	} else {
		sc.TraceOptions = trace.TraceOptions(0)
	}

	return &sc, true

}

func (c *RequestContext) Send(httpStatus int, v interface{}) {
	var body []byte
	defer func() {
		c.response.WriteHeader(httpStatus)
		if _, err := c.response.Write(body); err != nil {
			logger.Warn("Failed to write response", zap.Error(err))
		}
	}()

	body, ok := v.([]byte)
	if ok {
		return
	}

	body, err := json.Marshal(v)
	if err != nil {
		logger.Warn("Failed to marshal.", zap.Error(err))
		body = []byte(err.Error())
		return
	}
}

func (c *RequestContext) Response(httpStatus int, data interface{}) {
	code := http.StatusText(httpStatus)

	err, isError := data.(error)
	if isError {
		grpcStatus, ok := status.FromError(err)
		if !ok {
			data = err.Error()
		} else {
			data = grpcStatus.Message()
			code = grpcStatus.Code().String()
		}
	}

	c.Send(
		httpStatus,
		&ErrorResponse{
			Status:  httpStatus,
			Code:    code,
			Message: data,
		})
}

type ErrorResponse struct {
	Status  int
	Code    string
	Message interface{}
}
