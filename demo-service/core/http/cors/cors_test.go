package cors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// Testing constants.
//
const (
	validURL           = "http://api.virgilsecurity.com/cards/v5"
	validHeader        = "header"
	anotherValidHeader = "another header"
	validOriginHeader  = "*"
)

//
// WrapHTTPHandler :: for an OptionsGetter without values :: doesn't set any CORS headers.
//
func TestWrapHTTPHandlerForAnEmptyOptionsGetter(t *testing.T) {
	resp := httptest.NewRecorder()
	req := getRequestStub()
	handler := new(handlerStub)
	corsOpts := new(optsGetterStub)

	h := WrapHTTPHandler(handler, corsOpts)
	h.ServeHTTP(resp, req)

	// Asserts
	allowOrigin, ok := resp.HeaderMap[headerAccessControlAllowedOrigins]

	assert.False(t, ok)
	assert.Empty(t, allowOrigin)

	allowMethods, ok := resp.HeaderMap[headerAccessControlAllowedMethods]

	assert.False(t, ok)
	assert.Empty(t, allowMethods)

	allowHeaders, ok := resp.HeaderMap[headerAccessControlAllowedHeaders]

	assert.False(t, ok)
	assert.Empty(t, allowHeaders)

	allowCredentials, ok := resp.HeaderMap[headerAccessControlAllowedCredentials]

	assert.False(t, ok)
	assert.Empty(t, allowCredentials)

	exposeHeaders, ok := resp.HeaderMap[headerAccessControlExposedHeaders]

	assert.False(t, ok)
	assert.Empty(t, exposeHeaders)
}

//
// WrapHTTPHandler :: for an OptionsGetter with not empty origins values :: returns only Origin CORS headers.
//
func TestWrapForANotEmptyOrigin(t *testing.T) {
	resp := httptest.NewRecorder()
	req := getRequestStub()
	handler := new(handlerStub)
	corsOpts := new(optsGetterStub)
	corsOpts.allowedOrigins = validOriginHeader
	corsOpts.allowedMethods = http.MethodOptions
	req.Header.Add("Access-Control-Request-Method", "OPTIONS")
	req.Header.Add("Origin", "*")

	h := WrapHTTPHandler(handler, corsOpts)
	h.ServeHTTP(resp, req)

	// Asserts
	allowOrigin, ok := resp.HeaderMap[headerAccessControlAllowedOrigins]
	assert.True(t, ok)
	assert.Equal(t, pluralizeHeader(validOriginHeader), allowOrigin)

	allowMethods, ok := resp.HeaderMap[headerAccessControlAllowedMethods]
	assert.True(t, ok)
	assert.Equal(t, pluralizeHeader(http.MethodOptions), allowMethods)

	allowHeaders, ok := resp.HeaderMap[headerAccessControlAllowedHeaders]
	assert.False(t, ok)
	assert.Empty(t, allowHeaders)

	allowCredentials, ok := resp.HeaderMap[headerAccessControlAllowedCredentials]
	assert.False(t, ok)
	assert.Empty(t, allowCredentials)

	exposeHeaders, ok := resp.HeaderMap[headerAccessControlExposedHeaders]
	assert.False(t, ok)
	assert.Empty(t, exposeHeaders)
}

//
// normalizeHeader :: for an empty string :: returns an empty list.
//
func TestNormalizeHeaderForAnEmptyString(t *testing.T) {

	headers := normalizeHeader("")

	assert.Len(t, headers, 0)
}

//
// normalizeHeader :: for a single comma character :: returns an empty list.
//
func TestNormalizeHeaderForComma(t *testing.T) {

	headers := normalizeHeader(",")

	assert.Len(t, headers, 0)
}

//
// normalizeHeader :: for several commas :: returns an empty list.
//
func TestNormalizeHeaderForSeveralCommas(t *testing.T) {

	headers := normalizeHeader(",,,")

	assert.Len(t, headers, 0)
}

//
// normalizeHeader :: for a single value :: returns one entry.
//
func TestNormalizeHeaderForSingleValue(t *testing.T) {

	headers := normalizeHeader(validHeader)

	assert.Len(t, headers, 1)
	assert.Equal(t, headers[0], validHeader)
}

//
// normalizeHeader :: for a single value surrounded by commas :: returns one entry.
//
func TestNormalizeHeaderForSingleValueSurroundedByCommas(t *testing.T) {

	headers := normalizeHeader(fmt.Sprintf(" , ,  %s ,, ,", validHeader))

	assert.Len(t, headers, 1)
	assert.Equal(t, headers[0], validHeader)
}

//
// normalizeHeader :: for a several values surrounded by commas :: returns several entries.
//
func TestNormalizeHeaderForSeveralValuesSurroundedByCommas(t *testing.T) {

	headers := normalizeHeader(fmt.Sprintf(" , ,  %s ,, ,%s", validHeader, anotherValidHeader))

	assert.Len(t, headers, 2)
	assert.Equal(t, headers[0], validHeader)
	assert.Equal(t, headers[1], anotherValidHeader)
}

//
// getRequestStub returns an http request stub object.
//
func getRequestStub() *http.Request {
	req, err := http.NewRequest(http.MethodOptions, validURL, strings.NewReader(""))
	if nil != err {
		panic(err)
	}

	return req
}

//
// pluralizeHeader pluralizes header into headers slice.
//
func pluralizeHeader(h string) []string {

	return []string{h}
}
