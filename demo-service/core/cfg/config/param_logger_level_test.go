package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/log"
)

//
// NewLogLevel : initialize logger parameter with all available parameters : expected no validation errors.
//
func TestNewLogLevelShouldReturnNoErrorAndValidLoggerParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		DefaultValue = log.LevelError
		LoggerValue  = log.LevelError
	)

	loggerInfoParameter := NewLoggerLevel(EnvVariable, UsageMessage)
	loggerInfoParameter.level = LoggerValue

	err := loggerInfoParameter.validate()

	assert.Nil(t, err)
	assert.NotEqual(t, nil, loggerInfoParameter.getValuePointer())
	assert.Equal(t, DefaultValue, loggerInfoParameter.getDefault())
	assert.Equal(t, EnvVariable, loggerInfoParameter.getName())
	assert.Equal(t, UsageMessage, loggerInfoParameter.getUsage())
	assert.Equal(t, LoggerValue, loggerInfoParameter.getLevel())
	assert.Equal(t, ParameterTypeLogger, loggerInfoParameter.getType())
}

//
// NewLogLevel : initialize logger with incorrect log level : should return an error.
//
func TestNewLogLevelWithIncorrectLogLevelShouldReturnAnError(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
		LoggerValue  = "not supported log level"
	)

	loggerInfoParameter := NewLoggerLevel(EnvVariable, UsageMessage)
	loggerInfoParameter.level = LoggerValue

	assert.Error(t, loggerInfoParameter.validate())
}

//
// NewLogger : initialize logger without log level : should return an error.
//
func TestNewLogLevelParameterWithoutLogLevelShouldReturnAnError(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"
	)

	loggerInfoParameter := NewLoggerLevel(EnvVariable, UsageMessage)

	assert.Error(t, loggerInfoParameter.validate())
}
