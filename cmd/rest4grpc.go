package main

import (
	"flag"
	"github.com/tmarcus87/rest4grpc/grpc"
	"github.com/tmarcus87/rest4grpc/logger"
	"github.com/tmarcus87/rest4grpc/server"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strings"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	help = flags.Bool("help", false, prettify(`
			Print usage
			`))

	target = flags.String("target", "", prettify(`
			Specify the address of backend(gRPC) server.
			Available format is following.
			- grpc://127.0.0.1:5000
			`))

	bind = flags.String("bind", ":8888", prettify(`
			Specify the bind address of rest server.
			`))

	traceSamplingFraction = flags.Float64("trace-fraction", 0.1, prettify(`
			Specify the sampling fraction of distributed tracing.
			`))

	logLevel = flags.String("log-level", "info", prettify(`
			Specify log level [error|warn|info|debug]
			`))

	logEncoder = flags.String("log-encoder", "json", prettify(`
			Log encoder type.
			Available type is following.
			- json
			- console
			`))
)

var (
	s *http.Server
)

const (
	exitCodeRuntimeError = 1
	exitCodeInvalidParam = 3
)

func prettify(docString string) string {
	lines := strings.Split(docString, "\n")

	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 && line == "" {
			continue
		}
		lines[i] = line
		i++
	}

	return strings.Join(lines[:i], "\n")
}

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		logger.Error("Failed to parse arguments", zap.Error(err))
		os.Exit(exitCodeRuntimeError)
	}

	if err := logger.Setup(*logLevel, *logEncoder); err != nil {
		logger.Warn("Failed to setup logger", zap.Error(err))
	}

	if *help {
		flags.Usage()
		os.Exit(0)
	}

	if *target == "" {
		logger.Error("The -target arguments must be specified.")
		os.Exit(exitCodeInvalidParam)
	}

	client, err := grpc.NewDynamicGrpcClient(*target)
	if err != nil {
		logger.Error("failed to create grpc client", zap.Error(err))
		os.Exit(exitCodeInvalidParam)
	}

	handler, err := server.NewProxyHandler(client, *traceSamplingFraction)
	if err != nil {
		logger.Error("Failed to create proxy handler.", zap.Error(err))
		os.Exit(exitCodeInvalidParam)
	}

	s := server.NewServer(*bind, handler)

	if err := s.Start(); err != nil {
		logger.Error("Failed to serve server.", zap.Error(err))
		os.Exit(exitCodeRuntimeError)
	}
}
