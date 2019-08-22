// +build functional

package deal

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/KWRI/demo-service/src/router"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/test/helpers"

	"github.com/KWRI/demo-service/test/fixtures"
)

//
// TestGetSummaryApplicationStatisticEndpoint tests GET /application/{application_id}/summary endpoint.
// Success result should be: {
//  	"cards": 201,
//  	"requests": 301
// }
//
func TestGetSummaryApplicationStatisticEndpoint(t *testing.T) {

	httpRequest, err := http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			strings.Replace(
				router.RouteGetSummaryStatistics,
				router.IDPlaceholder,
				fixtures.ApplicationID,
				1,
			),
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

	defer func() {
		assert.Nil(t, httpResponse.Body.Close())
	}()

	body, err := ioutil.ReadAll(httpResponse.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.JSONEq(t, fixtures.ApplicationSummaryResponse, string(body))
}
