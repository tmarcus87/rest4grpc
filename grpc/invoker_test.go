package grpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"github.com/tmarcus87/rest4grpc"
	"go.opencensus.io/trace"
	"strings"
	"testing"
)

func TestDynamicGrpcClient_Invoke(t *testing.T) {
	svr := rest4grpc.NewGrpcServer()
	if err := svr.Start(8889); err != nil {
		panic(err)
	}
	defer svr.Stop()

	client, err := NewDynamicGrpcClient("grpc://localhost:8889")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx, span := trace.StartSpan(context.Background(), "rest4grpc-test")
	defer span.End()

	traceID := span.SpanContext().TraceID.String()

	res, err := client.Invoke(ctx, "pb.TestService", "MethodA", strings.NewReader(`{"name": "name", "age": 20}`))
	if err != nil {
		panic(err)
	}

	fmt.Println("tid : " + traceID)
	fmt.Println("res : " + string(res))

	asserts := assert.New(t)
	asserts.Equal("name", gjson.Get(string(res), "name").String())
	asserts.Equal(int64(20), gjson.Get(string(res), "age").Int())
	asserts.Equal(traceID, gjson.Get(string(res), "trace_id").String())
}
