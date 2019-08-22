package di

import (
	"github.com/KWRI/demo-service/src/app"
	"github.com/sarulabs/di"

	"github.com/KWRI/demo-service/src/dao"
)

//
// Definition name.
//
const (
	DefDealController = "DealController"
	DefDealValidator  = "DealValidator"
	DefDealRepository = "DealRepository"
)

//
// initDealRepository dependency registrar.
//
func (c *Container) initDealRepository() error {

	return c.RegisterDependency(
		DefDealRepository,
		func(ctx di.Container) (interface{}, error) {
			return dao.NewDealRepository(
				c.GetLogger(),
				c.GetMySQLDatabase(),
			), nil
		},
		nil,
	)
}

//
// GetDealRepository dependency getter.
//
func (c *Container) GetDealRepository() dao.DealRepositoryProvider {
	return c.Container.Get(DefDealRepository).(dao.DealRepositoryProvider)
}

//
// initApplicationValidator dependency registrar.
//
func (c *Container) initApplicationValidator() error {

	return c.RegisterDependency(
		DefDealValidator,
		func(ctx di.Container) (interface{}, error) {
			return app.NewValidator(
				c.GetDealRepository(),
			), nil
		},
		nil,
	)
}

//
// GetDealValidator dependency getter.
//
func (c *Container) GetDealValidator() app.ValidatorProvider {
	return c.Container.Get(DefDealValidator).(app.ValidatorProvider)
}

//
// initDealController dependency registrar.
//
func (c *Container) initDealController() error {

	return c.RegisterDependency(
		DefDealController,
		func(ctx di.Container) (interface{}, error) {
			return app.NewDeal(
				c.GetDealValidator(),
				c.GetDealRepository(),
			), nil
		},
		nil,
	)
}

//
// GetDealController dependency getter.
//
func (c *Container) GetDealController() app.ControllerProvider {
	return c.Container.Get(DefDealController).(app.ControllerProvider)
}
