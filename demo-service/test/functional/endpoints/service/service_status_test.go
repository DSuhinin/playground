// +build functional

package service

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	kitHTTP "github.com/KWRI/demo-service/core/http"
	"github.com/KWRI/demo-service/core/test/helpers"
)

func TestServiceStatusEndpoint(t *testing.T) {

	httpRequest, err := http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			kitHTTP.RouteServiceStatus,
			nil,
		),
		nil,
	)

	assert.Nil(t, err)

	HTTPClient := &http.Client{
		Timeout: time.Second * 1,
	}

	httpResponse, err := HTTPClient.Do(httpRequest)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)

	defer func() {
		assert.Nil(t, httpResponse.Body.Close())
	}()

	_, err = ioutil.ReadAll(httpResponse.Body)
	assert.Nil(t, err)
}
