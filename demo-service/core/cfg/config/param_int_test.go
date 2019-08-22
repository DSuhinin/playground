package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewInt : initialize int parameter with all available parameters : expected no validation errors.
//
func TestNewIntShouldReturnNoErrorAndValidIntParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = 1
		IntValue     = 0
	)

	intParameter := NewInt(EnvVariable, UsageMessage, DefaultValue)

	err := intParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, DefaultValue, intParameter.getDefault())
	assert.Equal(t, EnvVariable, intParameter.getName())
	assert.Equal(t, UsageMessage, intParameter.getUsage())
	assert.Equal(t, ParameterTypeInteger, intParameter.getType())
	assert.Equal(t, IntValue, intParameter.getValue())
}
