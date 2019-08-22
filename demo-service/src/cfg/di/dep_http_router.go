package di

import (
	"github.com/sarulabs/di"

	"github.com/KWRI/demo-service/core/http"
	"github.com/KWRI/demo-service/src/router"
)

//
// Definition name.
//
const (
	DefHTTPRouter = "HTTPRouter"
)

//
// initHTTPRouter dependency registrar.
//
func (c *Container) initHTTPRouter() error {
	return c.RegisterDependency(
		DefHTTPRouter,
		func(ctx di.Container) (interface{}, error) {
			r := http.NewRouter(
				c.GetLogger(),
				c.GetConfig().GetMetricPrefix(),
				http.SetupHealthDependencyList(c.GetMySQLDatabase()),
			)

			router.InitDealsRouteList(
				r,
				c.GetDealController(),
			)

			return r, nil
		},
		nil,
	)
}

//
// GetHTTPRouter dependency getter.
//
func (c *Container) GetHTTPRouter() *http.Router {
	return c.Container.Get(DefHTTPRouter).(*http.Router)
}
