package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewFloat64 : initialize float64 parameter with all available parameters : expected no validation errors.
//
func TestNewFloat64ShouldReturnNoErrorAndValidFloat64Parameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = 0.02
		Float64Value = 0.0
	)

	float64Parameter := NewFloat64(EnvVariable, UsageMessage, DefaultValue)

	err := float64Parameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, DefaultValue, float64Parameter.getDefault())
	assert.Equal(t, EnvVariable, float64Parameter.getName())
	assert.Equal(t, UsageMessage, float64Parameter.getUsage())
	assert.Equal(t, ParameterTypeFloat64, float64Parameter.getType())
	assert.Equal(t, Float64Value, float64Parameter.getValue())
}
