package config

import "time"

//
// Getter provides a base interface to access to the different configuration variables types.
//
type Getter interface {

	//
	// IsParsed returns the parsing status of Config object.
	//
	IsParsed() bool

	//
	// GetString returns the string representation of ENV parameter.
	//
	GetString(p string) string

	//
	// GetBase64String returns the base64 string representation of ENV parameter.
	//
	GetBase64String(p string) []byte

	//
	// GetBool returns the bool representation of ENV parameter.
	//
	GetBool(p string) bool

	//
	// GetInt returns the integer representation of ENV parameter.
	//
	GetInt(p string) int

	//
	// GetFloat64 returns the float64 representation of ENV parameter.
	//
	GetFloat64(p string) float64

	//
	// GetDuration returns the time.Duration representation of ENV parameter.
	//
	GetDuration(p string) time.Duration

	//
	// GetCassandraConnectionInfo returns the parsed cassandra configuration.
	//
	GetCassandraConnectionInfo(p string) *CassandraInfoParam

	//
	// GetRedisConnectionInfo returns the parsed redis configuration.
	//
	GetRedisConnectionInfo(p string) *RedisInfoParam

	//
	// GetConsulConnectionInfo returns the parsed redis configuration.
	//
	GetConsulConnectionInfo(p string) *ConsulInfoParam

	//
	// GetLogLevel returns the parsed logger configuration.
	//
	GetLogLevel(p string) string
}
