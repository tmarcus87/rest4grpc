package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	if err := Setup("info", "json"); err != nil {
		panic(err)
	}
}

func config() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.Encoding = "json"
	cfg.Development = false
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	return cfg
}

func Setup(logLevel, logEncoder string) error {
	cfg := config()

	lvl := zap.AtomicLevel{}
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		return fmt.Errorf("invalid log level : %w", err)
	}
	cfg.Level = lvl

	cfg.Encoding = logEncoder

	var err error
	logger, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	Debug("Setup logger",
		zap.String("level", lvl.String()),
		zap.String("encoder", logEncoder))

	return nil
}

func FromContext(ctx context.Context) *zap.Logger {
	var tid, sid string
	if v := ctx.Value("TraceID"); v != nil {
		tid = v.(string)
	}
	if v := ctx.Value("SpanID"); v != nil {
		sid = v.(string)
	}
	return logger.With(zap.String("trace_id", tid), zap.String("span_id", sid))
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}
