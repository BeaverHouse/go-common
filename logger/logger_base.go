package logger

import "github.com/BeaverHouse/go-common/env"

// Field represents a key-value pair for structured logging
type Field struct {
	Key   string
	Value any
}

// F creates a new Field
func F(key string, value any) Field {
	return Field{Key: key, Value: value}
}

// Logger is the interface for logging
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Sync() error
}

// NewLogger returns the logger suited to the current environment: a structured
// JSON ZapLogger for deployed environments (development, production) and a
// human-readable SimpleLogger for local development. Selecting by environment
// rather than by runtime type (server vs CLI) lets the same code log
// consistently across API servers, MCP servers, and batch jobs.
func NewLogger() (Logger, error) {
	if env.IsGoEnv(env.LocalEnv) {
		return NewSimpleLogger(), nil
	}
	return NewZapLogger()
}
