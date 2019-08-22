package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewStringBase64 : initialize string_base64 parameter with all available parameters : expected no validation errors.
//
func TestNewStringBase64ShouldReturnNoErrorAndValidBase64StringParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = "ZGVmYXVsdA=="
	)

	var (
		Base64StringValue = []byte{}
	)

	base64StringParameter := NewBase64String(EnvVariable, UsageMessage, DefaultValue)

	err := base64StringParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, DefaultValue, base64StringParameter.getDefault())
	assert.Equal(t, EnvVariable, base64StringParameter.getName())
	assert.Equal(t, UsageMessage, base64StringParameter.getUsage())
	assert.Equal(t, ParameterTypeBase64String, base64StringParameter.getType())
	assert.Equal(t, Base64StringValue, base64StringParameter.getValue())
}

//
// NewStringBase64 : initialize string_base64 parameter with incorrect base64 value : should return an error.
//
func TestNewStringBase64WithIncorrectBase64ValueShouldReturnAnError(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = "ZGVmYXVsdA=="
	)

	base64StringParameter := NewBase64String(EnvVariable, UsageMessage, DefaultValue)
	base64StringParameter.value = "incorrect base64 string"

	err := base64StringParameter.validate()

	assert.Error(t, err)
}
