// +build integration

package deal

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/KWRI/demo-service/src/router"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/test/helpers"

	"github.com/KWRI/demo-service/test/fixtures"
	"github.com/KWRI/demo-service/test/integration"
)

//
// TestGetDealEndpoint tests GET /deals/{deal_id} endpoint.
// Success result should be: {
//  	"deal_id": 888,
//  	"opportunity_id": "",
//  	"contract_id": 0,
//  	"kw_uid": "556396",
//  	"kw_uid_name": "Nicole Burton",
//  	"mc_id": 2,
//  	"mc_key": 0,
//  	"checklist_ids": "",
//  	"deal_owner": "556396",
//  	"deal_owner_name": "Nicole Burton",
//  	"deal_name": "Buyer Deal"
// }
//
func TestGetDealEndpoint(t *testing.T) {

	_, di := integration.TestUP(t)

	dealID := "888"

	httpRequest, err := http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			di.GetConfig().GetServerHTTPAddress(),
			strings.Replace(
				router.RouteGetDeal,
				router.IDPlaceholder,
				dealID,
				1,
			),
			nil,
		),
		nil,
	)
	assert.Nil(t, err)

	httpResponseRecorder := httptest.NewRecorder()
	di.GetHTTPRouter().GetMuxRouter().ServeHTTP(httpResponseRecorder, httpRequest)
	httpResponse := httpResponseRecorder.Result()

	body, err := ioutil.ReadAll(httpResponse.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.JSONEq(t, fixtures.DealGetResponse, string(body))
}

//
// TestGetDealEndpoint tests GET /deals endpoint.
// Success result should be: [
//  {
//    "deal_id": 761,
//    "opportunity_id": "O0-761",
//    "contract_id": 0,
//    "kw_uid": "556396",
//    "kw_uid_name": "Nicole Burton",
//    "mc_id": 2,
//    "mc_key": 0,
//    "checklist_ids": "",
//    "deal_owner": "615826",
//    "deal_owner_name": "",
//    "deal_name": "Buyer Deal"
//  },
//  {
//    "deal_id": 762,
//    "opportunity_id": "O0-762",
//    "contract_id": 0,
//    "kw_uid": "556396",
//    "kw_uid_name": "Nicole Burton",
//    "mc_id": 2,
//    "mc_key": 0,
//    "checklist_ids": "",
//    "deal_owner": "556396",
//    "deal_owner_name": "Nicole Burton",
//    "deal_name": "Buyer Deal"
//  },
//  {
//    "deal_id": 763,
//    "opportunity_id": "O0-763",
//    "contract_id": 0,
//    "kw_uid": "556396",
//    "kw_uid_name": "Nicole Burton",
//    "mc_id": 2,
//    "mc_key": 0,
//    "checklist_ids": "",
//    "deal_owner": "556396",
//    "deal_owner_name": "Nicole Burton",
//    "deal_name": "Buyer Deal"
//  }
//]
//
func TestGetDealListEndpoint(t *testing.T) {

	_, di := integration.TestUP(t)

	httpRequest, err := http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			di.GetConfig().GetServerHTTPAddress(),
			router.RouteGetDealsList,
			nil,
		),
		nil,
	)
	assert.Nil(t, err)

	httpResponseRecorder := httptest.NewRecorder()
	di.GetHTTPRouter().GetMuxRouter().ServeHTTP(httpResponseRecorder, httpRequest)
	httpResponse := httpResponseRecorder.Result()

	body, err := ioutil.ReadAll(httpResponse.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.JSONEq(t, fixtures.DealsListResponse, string(body))
}
