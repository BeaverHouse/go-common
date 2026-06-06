package logger

import (
	"fmt"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorGray   = "\033[90m"
)

// SimpleLogger writes human-readable, colorized lines, suited to local
// development in a terminal. Prefer NewLogger, which selects the format by
// environment; construct this directly only when console output is required.
type SimpleLogger struct{}

// NewSimpleLogger creates a SimpleLogger (human-readable console output).
func NewSimpleLogger() Logger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) formatFields(fields []Field) string {
	if len(fields) == 0 {
		return ""
	}
	var result strings.Builder
	for _, f := range fields {
		fmt.Fprintf(&result, " %s=%v", f.Key, f.Value)
	}
	return result.String()
}

func (l *SimpleLogger) log(level, color, msg string, fields ...Field) {
	fmt.Printf("%s%-5s%s %s%s\n", color, level, colorReset, msg, l.formatFields(fields))
}

func (l *SimpleLogger) Info(msg string, fields ...Field) {
	l.log("INFO", colorGreen, msg, fields...)
}

func (l *SimpleLogger) Error(msg string, fields ...Field) {
	l.log("ERROR", colorRed, msg, fields...)
}

func (l *SimpleLogger) Debug(msg string, fields ...Field) {
	l.log("DEBUG", colorGray, msg, fields...)
}

func (l *SimpleLogger) Warn(msg string, fields ...Field) {
	l.log("WARN", colorYellow, msg, fields...)
}

func (l *SimpleLogger) Sync() error {
	return nil
}
