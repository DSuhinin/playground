package stubs

import (
	"fmt"

	"github.com/stretchr/testify/mock"

	"github.com/KWRI/demo-service/core/log"
)

//
// Logger a stub that implements the log.Logger interface.
//
type Logger struct {
	mock.Mock
	data string
}

//
// Error logs a message with error severity.
//
func (l *Logger) Error(format string, args ...interface{}) {
	l.write(log.LevelError, format, args...)
}

//
// Debug logs a message with debug severity.
//
func (l *Logger) Debug(format string, args ...interface{}) {
	l.write(log.LevelDebug, format, args...)
}

//
// Info logs a message with info severity.
//
func (l *Logger) Info(format string, args ...interface{}) {
	l.write(log.LevelInfo, format, args...)
}

//
// Warn logs a message with info severity.
//
func (l *Logger) Warn(format string, args ...interface{}) {
	l.write(log.LevelWarn, format, args...)
}

//
// Health logs a message with health severity.
//
func (l *Logger) Health(format string, args ...interface{}) {
	l.write(log.LevelHealth, format, args...)
}

//
// write writes the string to the output.
//
func (l *Logger) write(level string, format string, args ...interface{}) {
	l.data = fmt.Sprintf("[%s] %s\n", level, fmt.Sprintf(format, args...))
}

//
// GetData returns logged data.
//
func (l *Logger) GetData() string {
	return l.data
}

//
// GetLogLevel returns the current Log Level.
//
func (l *Logger) GetLogLevel() string {
	return ""
}
