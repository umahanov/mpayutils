package log

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType string

const loggerKey loggerKeyType = "ctx-logger-fields"

type Logger struct {
	*zap.Logger
	zap.AtomicLevel
}

var logger Logger

func init() {
	level := zap.InfoLevel
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	}

	atomicLevel := zap.NewAtomicLevelAt(level)

	zapLogger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(os.Stderr),
			atomicLevel,
		),
	)

	logger = Logger{
		zapLogger,
		atomicLevel,
	}

	Info(context.Background(), "logger initialized", zap.String("log_level", level.String()))
}

func With(ctx context.Context, args ...zap.Field) context.Context {
	curFieldsAny := ctx.Value(loggerKey)
	if curFieldsAny == nil {
		return context.WithValue(ctx, loggerKey, args)
	}

	curFields, ok := curFieldsAny.([]zap.Field)
	if !ok {
		Error(context.Background(), "logger field in context has invalid type")
		return context.WithValue(ctx, loggerKey, args)
	}

	newFields := make([]zap.Field, len(curFields), len(curFields)+len(args))
	copy(newFields, curFields)
	newFields = append(newFields, args...)

	return context.WithValue(ctx, loggerKey, newFields)
}

func fields(ctx context.Context) []zap.Field {
	fieldsAny := ctx.Value(loggerKey)
	if fieldsAny == nil {
		return []zap.Field{}
	}

	fields, ok := fieldsAny.([]zap.Field)
	if !ok {
		Error(context.Background(), "logger field in context has invalid type")
		return []zap.Field{}
	}
	return fields
}

func Debug(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Debug(msg, fields...)
}

func Debugf(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Debug(fmt.Sprintf(format, args...), fields...)
}

func Info(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Info(msg, fields...)
}

func Infof(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Info(fmt.Sprintf(format, args...), fields...)
}

func Warn(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Warn(msg, fields...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Warn(fmt.Sprintf(format, args...), fields...)
}

func Error(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Error(msg, fields...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Error(fmt.Sprintf(format, args...), fields...)
}

func Panic(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Panic(msg, fields...)
}

func Panicf(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Panic(fmt.Sprintf(format, args...), fields...)
}

func Fatal(ctx context.Context, msg string, args ...zap.Field) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields)+len(args))
	copy(fields, ctxFields)
	fields = append(fields, args...)

	logger.Fatal(msg, fields...)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	ctxFields := fields(ctx)
	fields := make([]zap.Field, len(ctxFields), len(ctxFields))
	copy(fields, ctxFields)

	logger.Fatal(fmt.Sprintf(format, args...), fields...)
}

// Handler returns http handler to view/change log level in runtime,
// see go.uber.org/zap@v1.18.1/http_handler.go for usage
func Handler() http.Handler {
	return logger.AtomicLevel
}

func SetLevel(level string) error {
	switch strings.ToLower(level) {
	case "debug":
		logger.SetLevel(zap.DebugLevel)
	case "info":
		logger.SetLevel(zap.InfoLevel)
	case "warn":
		logger.SetLevel(zap.WarnLevel)
	case "error":
		logger.SetLevel(zap.ErrorLevel)
	case "fatal":
		logger.SetLevel(zap.FatalLevel)
	default:
		return fmt.Errorf("unkown log level: %s", level)
	}

	return nil
}
