package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewString : initialize duration parameter with all available parameters : expected no validation errors.
//
func TestNewStringShouldReturnNoErrorAndValidStringParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = "default string"
		StringValue  = "string"
	)

	stringParameter := NewString(EnvVariable, UsageMessage, DefaultValue)
	stringParameter.value = StringValue

	assert.Nil(t, stringParameter.validate())
	assert.Equal(t, DefaultValue, stringParameter.getDefault())
	assert.Equal(t, EnvVariable, stringParameter.getName())
	assert.Equal(t, UsageMessage, stringParameter.getUsage())
	assert.Equal(t, ParameterTypeString, stringParameter.getType())
	assert.Equal(t, StringValue, stringParameter.getValue())
}
