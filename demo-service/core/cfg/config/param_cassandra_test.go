package config

import (
	"fmt"
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

//
// Test Cassandra connection string parameters.
//
const (
	TestEnvVariable  = "ENV_VARIABLE"
	TestUsageMessage = "usage message"

	CassandraScheme      = "cassandra"
	CassandraUser        = "user"
	CassandraPassword    = "password"
	CassandraHost        = "host:9160"
	CassandraKeyspace    = "keyspace"
	CassandraDC          = "DC1"
	CassandraConsistency = "Quorum"
)

//
// TestNewCassandraInfoShouldReturnNoErrorAndValidCassandraInfoParameter ::
// initialize cassandra parameter with all available parameters : expected no validation errors.
//
func TestNewCassandraInfoShouldReturnNoErrorAndValidCassandraInfoParameter(t *testing.T) {

	var (
		ConnectionString = fmt.Sprintf(
			"%s://%s:%s@%s/%s?dc=%s&cl=%s",
			CassandraScheme,
			CassandraUser,
			CassandraPassword,
			CassandraHost,
			CassandraKeyspace,
			CassandraDC,
			CassandraConsistency,
		)
	)

	cassandraInfoParameter := NewCassandraInfo(TestEnvVariable, TestUsageMessage)
	cassandraInfoParameter.value = ConnectionString

	err := cassandraInfoParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, TestEnvVariable, cassandraInfoParameter.getName())
	assert.Equal(t, TestUsageMessage, cassandraInfoParameter.getUsage())
	assert.Equal(t, ParameterTypeCassandra, cassandraInfoParameter.getType())
	assert.Equal(t, ConnectionString, cassandraInfoParameter.GetConnectionString())
	assert.Equal(t, CassandraUser, cassandraInfoParameter.GetUser())
	assert.Equal(t, CassandraPassword, cassandraInfoParameter.GetPassword())
	assert.Equal(t, CassandraKeyspace, cassandraInfoParameter.GetKeyspace())
	assert.Equal(t, CassandraDC, cassandraInfoParameter.GetDataCenter())
	assert.Equal(t, []string{CassandraHost}, cassandraInfoParameter.GetHosts())
	assert.Equal(t, gocql.ParseConsistency(CassandraConsistency), cassandraInfoParameter.GetConsistencyLevel())
}

//
// TestDefaultConsistency :: expect cassandra consistency default value when it is not specified in connection string.
//
func TestDefaultConsistency(t *testing.T) {
	var (
		ConnectionString = fmt.Sprintf(
			"%s://%s:%s@%s/%s?dc=%s",
			CassandraScheme,
			CassandraUser,
			CassandraPassword,
			CassandraHost,
			CassandraKeyspace,
			CassandraDC,
		)
	)

	cassandraInfoParameter := NewCassandraInfo(TestEnvVariable, TestUsageMessage)
	cassandraInfoParameter.value = ConnectionString

	err := cassandraInfoParameter.validate()

	assert.Nil(t, err)
	assert.Equal(t, CassandraDefaultConsistencyLevel, cassandraInfoParameter.GetConsistencyLevel())
}
