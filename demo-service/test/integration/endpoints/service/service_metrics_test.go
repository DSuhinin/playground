// +build integration

package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	kitHTTP "github.com/KWRI/demo-service/core/http"
	"github.com/KWRI/demo-service/core/test/helpers"

	"github.com/KWRI/demo-service/test/integration"
)

func TestServiceMetricsEndpoint(t *testing.T) {

	_, di := integration.TestUP(t)

	httpRequest := httptest.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			di.GetConfig().GetServerHTTPAddress(),
			kitHTTP.RouteServiceMetrics,
			nil,
		),
		nil,
	)

	httpResponseRecorder := httptest.NewRecorder()
	di.GetHTTPRouter().GetMuxRouter().ServeHTTP(httpResponseRecorder, httpRequest)
	httpResponse := httpResponseRecorder.Result()

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
}
