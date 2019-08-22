package log

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/KWRI/demo-service/core/errors"
)

const (
	//
	// LevelHealth HEALTH logger level.
	//
	LevelHealth = "HEALTH"

	//
	// LevelDebug DEBUG logger level.
	//
	LevelDebug = "DEBUG"

	//
	// LevelInfo INFO logger level.
	//
	LevelInfo = "INFO"

	//
	// LevelWarn WARN logger level.
	//
	LevelWarn = "WARN"

	//
	// LevelError ERROR logger level.
	//
	LevelError = "ERROR"

	//
	// LevelHealthWeight weight of HEALTH logger level.
	//
	LevelHealthWeight = iota

	//
	// LevelDebugWeight weight of DEBUG logger level.
	//
	LevelDebugWeight

	//
	// LevelInfoWeight weight of INFO logger level.
	//
	LevelInfoWeight

	//
	// LevelWarnWeight weight of WARN logger level.
	//
	LevelWarnWeight

	//
	// LevelErrorWeight weight of ERROR logger level.
	//
	LevelErrorWeight
)

var (
	//
	// LevelList supported log level list.
	//
	LevelList = map[string]int{
		LevelHealth: LevelHealthWeight,
		LevelDebug:  LevelDebugWeight,
		LevelInfo:   LevelInfoWeight,
		LevelWarn:   LevelWarnWeight,
		LevelError:  LevelErrorWeight,
	}
)

//
// Log represents the logger.
//
type Log struct {
	mu       sync.RWMutex
	writer   io.Writer
	logLevel string
}

//
// New returns an instance of a logger.
//
func New(writer io.Writer, logLevel string) *Log {
	return &Log{
		writer:   writer,
		logLevel: logLevel,
	}
}

//
// Error logs error info.
//
func (l *Log) Error(format string, args ...interface{}) {
	l.write(LevelError, format, args...)
}

//
// Info logs info.
//
func (l *Log) Info(format string, args ...interface{}) {
	l.write(LevelInfo, format, args...)
}

//
// Warn logs info.
//
func (l *Log) Warn(format string, args ...interface{}) {
	l.write(LevelWarn, format, args...)
}

//
// Debug logs debug info.
//
func (l *Log) Debug(format string, args ...interface{}) {
	l.write(LevelDebug, format, args...)
}

//
// Health logs health info.
//
func (l *Log) Health(format string, args ...interface{}) {
	l.write(LevelHealth, format, args...)
}

//
// GetLogLevel returns the current Log Level.
//
func (l *Log) GetLogLevel() string {
	return l.logLevel
}

//
// write writes the string to the output.
//
func (l *Log) write(level string, format string, args ...interface{}) {

	if level != LevelHealth && level != LevelWarn && level != LevelInfo {
		if LevelList[strings.ToUpper(level)] < LevelList[strings.ToUpper(l.logLevel)] {
			return
		}
	}

	l.mu.RLock()
	{
		if _, err := l.writer.Write([]byte(
			fmt.Sprintf("[%s] %s\n", level, fmt.Sprintf(format, args...)),
		)); err != nil {
			panic(errors.WithMessage(err, `error while writing an log message`))
		}
	}
	l.mu.RUnlock()

}
