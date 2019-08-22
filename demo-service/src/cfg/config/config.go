package config

import (
	"time"

	"github.com/KWRI/demo-service/core/cfg/config"
)

//
// Common service variables.
//
const (
	MetricPrefix = "deals_demo_service"
	ServiceName  = "kwri-deals-service"
)

//
// Service ENV parameter list.
//
const (
	ConfMySQLDSN           = "DEMO_SERVICE_MYSQL_DSN"
	ConfLogLevel           = "DEMO_SERVICE_LOG_LEVEL"
	ConfServerAddress      = "DEMO_SERVICE_SERVER_ADDRESS"
	ConfServerReadTimeout  = "DEMO_SERVICE_SERVER_READ_TIMEOUT"
	ConfServerWriteTimeout = "DEMO_SERVICE_SERVER_WRITE_TIMEOUT"
)

//
// Config main config structure which stores predefined flags, env vars, etc...
//
type Config struct {
	*config.Config
}

//
// New returns Parameter storage with all already parsed params.
//
func New() (*Config, error) {

	c := config.New()
	c.RegisterParameters(
		config.NewString(
			ConfMySQLDSN,
			"MySQL connection DSN",
			"",
		),

		config.NewLoggerLevel(
			ConfLogLevel,
			"Service log level",
		),

		config.NewString(
			ConfServerAddress,
			"Specify address to listen HTTP requests.",
			":8080",
		),
		config.NewDuration(
			ConfServerReadTimeout,
			"Service read timeout", 3*time.Second,
		),
		config.NewDuration(
			ConfServerWriteTimeout,
			"Service write timeout", 5*time.Second,
		),
	)

	if err := c.Parse(); err != nil {
		return nil, err
	}

	return &Config{
		Config: c,
	}, nil
}
