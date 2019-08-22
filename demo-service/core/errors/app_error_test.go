package errors

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Code = 123
	Msg  = "error-msg"
	JSON = fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", Code, Msg)
)

//
// TestAppError :: check is error code and message are the same as specified.
//
func TestAppError(t *testing.T) {
	appError := NewAppError(Code, Msg)

	assert.Equal(t, Code, appError.GetCode())
	assert.Equal(t, Msg, appError.GetMessage())
}

//
// TestAppError_ToJSON :: check correct error object serialization to JSON string.
//
func TestAppError_ToJSON(t *testing.T) {
	errorJson, err := json.Marshal(NewAppError(Code, Msg))

	assert.NoError(t, err)
	assert.NotEmpty(t, errorJson)
	assert.Equal(t, JSON, string(errorJson))
}

//
// TestAppError_FromJSON :: check correct error object deserialization from JSON string.
//
func TestAppError_FromJSON(t *testing.T) {
	appError := &AppError{}
	err := json.Unmarshal([]byte(JSON), appError)

	assert.NoError(t, err)
	assert.NotEmpty(t, appError)
	assert.Equal(t, Code, appError.GetCode())
	assert.Equal(t, Msg, appError.GetMessage())
}
