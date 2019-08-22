package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewBool : initialize bool parameter with all available parameters : expected no validation errors.
//
func TestNewBoolShouldReturnNoErrorAndValidBoolParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = true
		BoolValue    = false
	)

	boolParameter := NewBool(EnvVariable, UsageMessage, DefaultValue)

	err := boolParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, DefaultValue, boolParameter.getDefault())
	assert.Equal(t, EnvVariable, boolParameter.getName())
	assert.Equal(t, UsageMessage, boolParameter.getUsage())
	assert.Equal(t, ParameterTypeBool, boolParameter.getType())
	assert.Equal(t, BoolValue, boolParameter.getValue())
}
