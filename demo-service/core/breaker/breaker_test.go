package breaker

import (
	"testing"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/test/helpers"
)

//
// TestInit :: ensure that config commands passed into Init method are stored in Hystrix circuit brakes settings.
//
func TestInit(t *testing.T) {
	commandConfig := CommandConfig{}
	commandName := helpers.GenerateRandomString()
	settings := hystrix.GetCircuitSettings()

	assert.Empty(t, settings[commandName])

	commandsConfig := make(CommandConfigList)
	commandsConfig[commandName] = commandConfig
	Init(commandsConfig)
	settings = hystrix.GetCircuitSettings()

	assert.NotEmpty(t, settings[commandName])
}
