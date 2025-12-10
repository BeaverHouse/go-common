package logger

import (
	"fmt"
	"time"
)

// SimpleLogger is a simple fmt-based logger for CLI/batch use and testing
type SimpleLogger struct{}

// NewSimpleLogger creates a new SimpleLogger
func NewSimpleLogger() Logger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) formatFields(fields []Field) string {
	if len(fields) == 0 {
		return ""
	}
	result := ""
	for _, f := range fields {
		result += fmt.Sprintf(" %s=%v", f.Key, f.Value)
	}
	return result
}

func (l *SimpleLogger) log(level, msg string, fields ...Field) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s: %s%s\n", timestamp, level, msg, l.formatFields(fields))
}

func (l *SimpleLogger) Info(msg string, fields ...Field) {
	l.log("INFO", msg, fields...)
}

func (l *SimpleLogger) Error(msg string, fields ...Field) {
	l.log("ERROR", msg, fields...)
}

func (l *SimpleLogger) Debug(msg string, fields ...Field) {
	l.log("DEBUG", msg, fields...)
}

func (l *SimpleLogger) Warn(msg string, fields ...Field) {
	l.log("WARN", msg, fields...)
}

func (l *SimpleLogger) Sync() error {
	return nil
}
