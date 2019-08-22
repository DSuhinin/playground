package di

import (
	"github.com/KWRI/demo-service/core/cfg/di"

	"github.com/KWRI/demo-service/src/cfg/config"
)

//
// Container is a dependency resolver object.
//
type Container struct {
	config *config.Config
	*di.Container
}

//
// NewContainer returns an instance of the DI DIContainer.
//
func NewContainer(config *config.Config) (*Container, error) {

	container, err := di.NewContainer()
	if err != nil {
		return nil, err
	}

	return &Container{
		config:    config,
		Container: container,
	}, nil
}

//
// Build builds the application dependencies at once.
//
func (c *Container) Build() error {

	dependencies := []func() error{
		c.initLogger,
		c.initHTTPRouter,
		c.initDealRepository,
		c.initDealController,
		c.initApplicationValidator,
		c.initMySQL,
	}

	for _, dependency := range dependencies {
		if err := dependency(); err != nil {
			return err
		}
	}

	c.Container.Build()

	return nil
}
