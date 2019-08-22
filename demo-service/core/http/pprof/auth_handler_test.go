package pprof

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/test/helpers"
)

//
// TestAuthHandler_NoToken_Unauthorized :: check for Unauthorized response when no token provided.
//
func TestAuthHandler_NoToken_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost", nil)
	assert.NoError(t, err)

	writer := httptest.NewRecorder()
	handler := AuthHandler(helpers.GenerateRandomString())
	handler.ServeHTTP(writer, req)

	assert.Equal(t, http.StatusUnauthorized, writer.Result().StatusCode)
}

//
// TestAuthHandler_WrongToken_Unauthorized :: check for Unauthorized response when wrong token provided.
//
func TestAuthHandler_WrongToken_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost?token="+helpers.GenerateRandomString(), nil)
	assert.NoError(t, err)

	writer := httptest.NewRecorder()
	handler := AuthHandler(helpers.GenerateRandomString())
	handler.ServeHTTP(writer, req)

	assert.Equal(t, http.StatusUnauthorized, writer.Result().StatusCode)
}

//
// TestAuthHandler_CorrectToken_OK :: check for OK response when correct token provided.
//
func TestAuthHandler_CorrectToken_OK(t *testing.T) {
	token := helpers.GenerateRandomString()
	req, err := http.NewRequest("GET", "http://localhost?token="+token, nil)
	assert.NoError(t, err)

	writer := httptest.NewRecorder()
	handler := AuthHandler(token)
	handler.ServeHTTP(writer, req)

	assert.Equal(t, http.StatusOK, writer.Result().StatusCode)
}
