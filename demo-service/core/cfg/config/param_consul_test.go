package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// TestNewConsulInfoShouldReturnNoErrorAndValidConsulInfoParameter ::
// initialize cassandra parameter with all available parameters : expected no validation errors.
//
func TestNewConsulInfoShouldReturnNoErrorAndValidConsulInfoParameter(t *testing.T) {

	//
	// Test Cassandra connection string parameters.
	//
	const (
		TestEnvVariable  = "ENV_VARIABLE"
		TestUsageMessage = "usage message"

		ConsulScheme   = "consul"
		ConsulUser     = "user"
		ConsulPassword = "password"
		ConsulHost     = "host:9160"
		ConsulDC       = "DC1"
	)

	var (
		ConnectionString = fmt.Sprintf(
			"%s://%s:%s@%s?dc=%s",
			ConsulScheme,
			ConsulUser,
			ConsulPassword,
			ConsulHost,
			ConsulDC,
		)
	)

	consulInfoParameter := NewConsulInfo(TestEnvVariable, TestUsageMessage)
	consulInfoParameter.value = ConnectionString

	err := consulInfoParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, TestEnvVariable, consulInfoParameter.getName())
	assert.Equal(t, TestUsageMessage, consulInfoParameter.getUsage())
	assert.Equal(t, ParameterTypeConsul, consulInfoParameter.getType())
	assert.Equal(t, ConnectionString, consulInfoParameter.GetConnectionString())
	assert.Equal(t, ConsulUser, consulInfoParameter.GetUser())
	assert.Equal(t, ConsulPassword, consulInfoParameter.GetPassword())
	assert.Equal(t, ConsulDC, consulInfoParameter.GetDataCenter())
	assert.Equal(t, CassandraHost, consulInfoParameter.GetHost())
}
