package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/test/helpers"

	"github.com/KWRI/demo-service/src/cfg/config"
	"github.com/KWRI/demo-service/src/cfg/di"
)

//
// TestUP prepares all needed variables, logic for functional tests.
//
func TestUP(t *testing.T) (*config.Config, *di.Container) {

	// Prepare test.
	helpers.ResetEnvVariables()
	helpers.ResetPrometheusMetrics()

	c, err := config.New()
	assert.Nil(t, err)

	// Build DI container.
	di, err := di.NewContainer(c)
	assert.Nil(t, err)
	assert.Nil(t, di.Build())

	return c, di
}
