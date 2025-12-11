package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZapLogger(t *testing.T) {
	logger, err := NewZapLogger()
	assert.NoError(t, err)
	assert.NotNil(t, logger)
	var _ Logger = logger
}

func TestNewSimpleLogger(t *testing.T) {
	logger := NewSimpleLogger()
	assert.NotNil(t, logger)
	var _ Logger = logger
}

func TestField(t *testing.T) {
	f := F("key", "value")
	assert.Equal(t, "key", f.Key)
	assert.Equal(t, "value", f.Value)
}

func TestSimpleLoggerMethods(t *testing.T) {
	logger := NewSimpleLogger()
	logger.Info("info message")
	logger.Error("error message")
	logger.Debug("debug message")
	logger.Warn("warn message")
	logger.Info("message with fields", F("key1", "value1"), F("key2", 42))
	assert.NoError(t, logger.Sync())
}

func TestZapLoggerMethods(t *testing.T) {
	logger, err := NewZapLogger()
	assert.NoError(t, err)
	logger.Info("info message")
	logger.Error("error message")
	logger.Debug("debug message")
	logger.Warn("warn message")
	logger.Info("message with fields", F("key1", "value1"), F("key2", 42))
}

func TestZapLoggerGetUnderlying(t *testing.T) {
	logger, err := NewZapLogger()
	assert.NoError(t, err)
	assert.NotNil(t, logger.GetZapLogger())
}
