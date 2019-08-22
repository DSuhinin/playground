package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/breaker"
	"github.com/KWRI/demo-service/core/log"
)

const (
	DefaultTimeout = 3 * 1000
)

//
// TestNewClientWithoutCircuitBreaker :: expect error client init when no circuit breaker setup.
//
func TestNewClientWithoutCircuitBreaker(t *testing.T) {
	client, err := initClient()

	assert.Error(t, err)
	assert.Empty(t, client)
}

//
// TestNewClient :: check no error for normal client init.
//
func TestNewClient(t *testing.T) {
	initCircuitBreaker()

	client, err := initClient()

	assert.NoError(t, err)
	assert.NotEmpty(t, client)
}

//
// TestClient_Do :: ensure that Do method returns correct response, without errors.
//
func TestClient_Do(t *testing.T) {
	expectedBody := []byte("response body")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write(expectedBody)
	}))

	initCircuitBreaker()
	client, err := initClient()
	assert.NoError(t, err)

	response, err := client.Do(&ClientRequest{
		Method: http.MethodGet,
		Route:  testServer.URL,
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expectedBody, response.Body)
}

//
// TestClientRetryOnEOFError :: testing HTTP-client resilience on EOF error caused by server crash during reply.
// Making sure that client retries request and receives success response eventually.
//
func TestClientRetryOnEOFError(t *testing.T) {
	i := 0
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if i == 0 {
			i++
			panic("Emulate server crash while response")
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("response body"))
	}))

	initCircuitBreaker()
	client, err := initClient()
	assert.NoError(t, err)

	response, err := client.Do(&ClientRequest{
		Method: http.MethodGet,
		Route:  testServer.URL,
	})

	assert.NoError(t, err)
	assert.Equal(t, StatusOK, response.StatusCode)
	assert.NotEmpty(t, response)
}

//
// initClient :: init HTTP client.
//
func initClient() (*Client, error) {
	return NewClient(log.New(ioutil.Discard, "debug"))
}

//
// initCircuitBreaker :: initialization of circuit breaker for HTTP client.
//
func initCircuitBreaker() {
	breaker.Init(map[string]breaker.CommandConfig{
		breaker.CommandCoreHTTPClient: {
			Timeout: DefaultTimeout,
		},
	})
}
