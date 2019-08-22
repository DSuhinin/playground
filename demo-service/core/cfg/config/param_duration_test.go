package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//
// NewDuration : initialize duration parameter with all available parameters : expected no validation errors.
//
func TestNewDurationShouldReturnNoErrorAndValidDurationParameter(t *testing.T) {

	const (
		EnvVariable   = "ENV_VARIABLE"
		UsageMessage  = "usage message"
		DefaultValue  = 1 * time.Second
		DurationValue = 0 * time.Second
	)

	durationParameter := NewDuration(EnvVariable, UsageMessage, DefaultValue)

	err := durationParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, DefaultValue, durationParameter.getDefault())
	assert.Equal(t, EnvVariable, durationParameter.getName())
	assert.Equal(t, UsageMessage, durationParameter.getUsage())
	assert.Equal(t, ParameterTypeDuration, durationParameter.getType())
	assert.Equal(t, DurationValue, durationParameter.getValue())
}
