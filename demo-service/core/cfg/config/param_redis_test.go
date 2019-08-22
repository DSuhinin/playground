package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// NewRedisInfo : initialize redis parameter with all available parameters : expected no validation errors.
//
func TestNewRedisInfoShouldReturnNoErrorAndValidCassandraInfoParameter(t *testing.T) {

	const (
		EnvVariable  = "ENV_VARIABLE"
		UsageMessage = "usage message"

		CassandraScheme = "redis"
		CassandraHost   = "host"
		CassandraPort   = "2345"
	)

	var (
		connectionString = fmt.Sprintf(
			"%s://%s:%s",
			CassandraScheme,
			CassandraHost,
			CassandraPort,
		)
	)

	redisInfoParameter := NewRedisInfo(EnvVariable, UsageMessage)
	redisInfoParameter.value = connectionString

	err := redisInfoParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, EnvVariable, redisInfoParameter.getName())
	assert.Equal(t, UsageMessage, redisInfoParameter.getUsage())
	assert.Equal(t, ParameterTypeRedis, redisInfoParameter.getType())
	assert.Equal(t, CassandraHost, redisInfoParameter.GetHost())
	assert.Equal(t, CassandraPort, redisInfoParameter.GetPort())
}
