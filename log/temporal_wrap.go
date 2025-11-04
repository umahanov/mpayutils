package log

import (
	"fmt"

	"go.uber.org/zap"
)

type TemporalLoggerWrap struct {
	logger *Logger
}

func WithTemporalWrap() *TemporalLoggerWrap {
	return &TemporalLoggerWrap{
		logger: &logger,
	}
}

func (w *TemporalLoggerWrap) Debug(msg string, keyvals ...any) {
	fields := make([]zap.Field, 0, len(keyvals)/2)
	for i := 1; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i-1]), keyvals[i]))
	}

	w.logger.Debug(msg, fields...)
}

func (w *TemporalLoggerWrap) Info(msg string, keyvals ...any) {
	fields := make([]zap.Field, 0, len(keyvals)/2)
	for i := 1; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i-1]), keyvals[i]))
	}

	w.logger.Info(msg, fields...)
}
func (w *TemporalLoggerWrap) Warn(msg string, keyvals ...any) {
	fields := make([]zap.Field, 0, len(keyvals)/2)
	for i := 1; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i-1]), keyvals[i]))
	}

	w.logger.Warn(msg, fields...)
}
func (w *TemporalLoggerWrap) Error(msg string, keyvals ...any) {
	fields := make([]zap.Field, 0, len(keyvals)/2)
	for i := 1; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i-1]), keyvals[i]))
	}

	w.logger.Error(msg, fields...)
}
