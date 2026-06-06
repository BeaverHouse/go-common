package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger writes structured JSON logs, suited to deployed environments where
// logs are shipped to an aggregator. Prefer NewLogger, which selects the format
// by environment; construct this directly only when JSON output is required.
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger creates a ZapLogger (structured JSON output).
func NewZapLogger() (*ZapLogger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: zapLogger}, nil
}

func (l *ZapLogger) toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// GetZapLogger returns the underlying zap.Logger
func (l *ZapLogger) GetZapLogger() *zap.Logger {
	return l.logger
}
