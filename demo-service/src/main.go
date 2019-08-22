package main

import (
	"fmt"
	"os"

	"github.com/KWRI/demo-service/core/http"

	"github.com/KWRI/demo-service/core/log"

	"github.com/KWRI/demo-service/src/cfg/config"
	"github.com/KWRI/demo-service/src/cfg/di"
)

//
// main entry point.
//
func main() {

	// Config, Logger
	c := initConfig()
	l := initLogger(c.GetLoggerLevel())

	// DI
	diContainer := initDIContainer(c)

	// Run Service
	var h = diContainer.GetHTTPRouter().GetMuxRouter()
	http.NewService(
		config.ServiceName,
		c.GetServerHTTPAddress(),
		c.GetServerReadTimeout(),
		c.GetServerWriteTimeout(),
		h,
		l,
		diContainer,
	).Run()
}

//
// initConfig makes Config init.
//
func initConfig() *config.Config {

	c, err := config.New()
	if err != nil {
		panicError("config initialization error", err)
	}

	return c
}

//
// initLogger makes Logger init.
//
func initLogger(logLevel string) log.Logger {

	return log.New(os.Stdout, logLevel)
}

//
// initDIContainer makes DI Container init.
//
func initDIContainer(c *config.Config) *di.Container {

	diContainer, err := di.NewContainer(c)
	if nil != err {
		panicError("DI initialization error", err)
	}

	if err := diContainer.Build(); err != nil {
		panicError("DI build error", err)
	}

	return diContainer
}

//
// panicError panics with an error.
//
func panicError(msg string, err error) {
	panic(fmt.Sprintf("%s: %+v", msg, err))
}
