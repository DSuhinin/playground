package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/KWRI/demo-service/src/app"

	kitHTTP "github.com/KWRI/demo-service/core/http"
	"github.com/KWRI/demo-service/core/http/response"
)

//
// Available application deal routes.
//
const (
	IDParameterName   = "deal_id"
	IDPlaceholder     = "{" + IDParameterName + "}"
	RouteGetDealsList = "/deals"
	RouteGetDeal      = "/deals/" + IDPlaceholder
)

//
// InitDealsRouteList initialize routes to handle Deal actions.
//
func InitDealsRouteList(router kitHTTP.RouterProvider, dealsController app.ControllerProvider) {

	// GET /deals route.
	router.Get(RouteGetDealsList, func(req *http.Request) response.Provider {

		data, err := dealsController.GetList()
		if err != nil {
			return response.New(err)
		}

		return response.New(data)
	})

	// GET /deals/{deal_id} route.
	router.Get(RouteGetDeal, func(req *http.Request) response.Provider {

		data, err := dealsController.Get(mux.Vars(req)[IDParameterName])
		if err != nil {
			return response.New(err)
		}

		return response.New(data)
	})
}
