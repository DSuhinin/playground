package config

import (
	"github.com/KWRI/demo-service/core/errors"
)

//
// Cassandra errors.
//
var (
	ErrCassandraConnectionStringIsEmpty     = errors.New("cassandra connection string is empty")
	ErrCassandraConnectionStringIsIncorrect = errors.New("cassandra connection string is incorrect")
	ErrCassandraSchemeIsIncorrect           = errors.New("cassandra connection protocol is incorrect")
	ErrCassandraHostsAreEmpty               = errors.New("cassandra hosts are empty")
	ErrCassandraKeyspaceIsEmpty             = errors.New("cassandra cassandraKeyspace is empty")
	ErrCassandraPasswordIsEmpty             = errors.New("cassandra password cannot be empty if username was set")
	ErrCassandraConsistencyLevelIsIncorrect = errors.New("cassandra consistency level is incorrect")
)

//
// Consul errors.
//
var (
	ErrConsulConnectionStringIsEmpty     = errors.New("consul connection string is empty")
	ErrConsulConnectionStringIsIncorrect = errors.New("consul connection string is incorrect")
	ErrConsulSchemeIsIncorrect           = errors.New("consul connection protocol is incorrect")
	ErrConsulHostIsEmpty                 = errors.New("consul host is empty")
	ErrConsulPasswordIsEmpty             = errors.New("consul password cannot be empty if username was set")
)

//
// Redis errors.
//
var (
	ErrRedisConnectionStringIsEmpty     = errors.New("redis connection string is empty")
	ErrRedisConnectionStringIsIncorrect = errors.New("redis connection string is incorrect")
	ErrRedisProtocolIsIncorrect         = errors.New("redis connection protocol is incorrect")
	ErrRedisHostIsEmpty                 = errors.New("redis host is empty")
	ErrRedisPortIsEmpty                 = errors.New("redis port is empty")
)
