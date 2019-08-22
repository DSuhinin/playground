package di

import (
	"os"

	"github.com/sarulabs/di"

	"github.com/KWRI/demo-service/core/log"
)

//
// Definition name.
//
const (
	DefLogger = "Logger"
)

//
// initLogger dependency registrar.
//
func (c *Container) initLogger() error {

	return c.RegisterDependency(
		DefLogger,
		func(ctx di.Container) (interface{}, error) {
			return log.New(os.Stdout, c.config.GetLoggerLevel()), nil
		},
		nil,
	)
}

//
// GetLogger dependency getter.
//
func (c *Container) GetLogger() log.Logger {
	return c.Container.Get(DefLogger).(log.Logger)
}
