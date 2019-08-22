package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// writer implements the io.Writer interface.
//
type writer struct {
	data []byte
}

//
// Write writes len(p) bytes from p to the underlying data stream.
//
func (w *writer) Write(p []byte) (n int, err error) {
	w.data = p
	return 0, nil
}

//
// GetData returns the data which has written into the stream.
//
func (w *writer) GetData() string {
	return string(w.data)
}

//
// Log_Debug : log message with DEBUG level : expected appropriate log message with status DEBUG.
//
func TestLog_Debug(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelDebug)
	logger.Debug("message")

	assert.Equal(t, "[DEBUG] message\n", writer.GetData())
}

//
// Log_Error : log message with ERROR level : expected appropriate log message with status ERROR.
//
func TestLog_Error(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelError)
	logger.Error("message")

	assert.Equal(t, "[ERROR] message\n", writer.GetData())
}

//
// Log_Health : log message with HEALTH level : expected appropriate log message with status HEALTH.
//
func TestLog_Health(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelHealth)
	logger.Health("message")

	assert.Equal(t, "[HEALTH] message\n", writer.GetData())
}

//
// Log_Info : log message with INFO level : expected appropriate log message with status INFO.
//
func TestLog_Info(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelInfo)
	logger.Info("message")

	assert.Equal(t, "[INFO] message\n", writer.GetData())
}

//
// Log_Warn : log message with WARN level : expected appropriate log message with status WARN.
//
func TestLog_Warn(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelWarn)
	logger.Warn("message")

	assert.Equal(t, "[WARN] message\n", writer.GetData())
}

//
// Log : set the log level upper then HEALTH and log message with HEALTH level
// 	   : expected appropriate log message with status HEALTH.
//
func TestLog_SetLogLevelUpperThenHealthStatusShouldLogHealthMessage(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelError)
	logger.Health("message")

	assert.Equal(t, "[HEALTH] message\n", writer.GetData())
}

//
// Log : set the log level upper then WARN and log message with WARN level
// 	   : expected appropriate log message with status WARN.
//
func TestLog_SetLogLevelUpperThenWarnStatusShouldLogWarnMessage(t *testing.T) {

	writer := &writer{}

	logger := New(writer, LevelError)
	logger.Warn("message")

	assert.Equal(t, "[WARN] message\n", writer.GetData())
}
