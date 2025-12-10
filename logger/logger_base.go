package logger

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
