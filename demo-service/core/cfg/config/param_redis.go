package config

import (
	"net/url"
	"strings"

	"github.com/KWRI/demo-service/core/errors"
)

const (
	//
	// RedisConnectionProtocol redis schema.
	//
	RedisConnectionProtocol = "redis"
)

//
// RedisConnectionInfoProvider declares all the connection info getters.
//
type RedisConnectionInfoProvider interface {
	//
	// GetHost returns the Redis host.
	//
	GetHost() string

	//
	// GetPort returns the Redis port.
	//
	GetPort() string
}

//
// ParameterRedisProvider is an interface for the redis parameter entry.
//
type ParameterRedisProvider interface {
	//
	// ParameterProvider include base parameter interface.
	//
	ParameterProvider

	//
	// getValuePointer returns value link.
	//
	getValuePointer() *string
}

//
// RedisInfoParam is a redis cache connection info parameter.
//
type RedisInfoParam struct {
	host  string
	port  string
	value string

	baseParam
}

//
// NewRedisInfo returns a new instance of redis parameter object.
//
func NewRedisInfo(name, usage string) *RedisInfoParam {

	return &RedisInfoParam{
		baseParam: baseParam{
			name:      name,
			usage:     usage,
			paramType: ParameterTypeRedis,
		},
	}
}

//
// GetValueLink returns value link.
//
func (p *RedisInfoParam) getValuePointer() *string {

	return &p.value
}

//
// GetHost returns the redis host.
//
func (p *RedisInfoParam) GetHost() string {

	return p.host
}

//
// GetPort returns redis port.
//
func (p *RedisInfoParam) GetPort() string {

	return p.port
}

//
// validate validates the redis connection string parameter.
//
func (p *RedisInfoParam) validate() error {

	if p.value == "" {
		return ErrRedisConnectionStringIsEmpty
	}

	connectionInfo, err := url.Parse(p.value)
	if nil != err {
		return errors.Wrap(err, ErrRedisConnectionStringIsIncorrect.WithMessage(
			"redis connection info (%s) parsing error", p.value,
		))
	}

	if _, err = validateConnectionProtocolClause(connectionInfo, RedisConnectionProtocol); nil != err {
		return errors.Wrap(err, ErrRedisProtocolIsIncorrect.WithMessage(
			"redis connection protocol (%s) is not (%s)", connectionInfo.Scheme, RedisConnectionProtocol,
		))
	}

	if p.host, err = validateRedisHostClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "redis host (%s) is invalid", connectionInfo.Host)
	}

	if p.port, err = validateRedisPortClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "redis port (%s) is invalid", connectionInfo.Port())
	}

	return nil
}

//
// validateHostClause validates the hosts clause.
//
func validateRedisHostClause(url *url.URL) (string, error) {

	host := strings.Split(url.Host, ":")[0]
	if "" == host {
		return "", ErrRedisHostIsEmpty
	}

	return host, nil
}

//
// validateRedisPortClause validates the hosts clause.
//
func validateRedisPortClause(url *url.URL) (string, error) {

	port := url.Port()
	if "" == port {
		return "", ErrRedisPortIsEmpty
	}

	return port, nil
}

//
// validateCassandraConnectionProtocolClause validates the connection protocol clause.
//
func validateConnectionProtocolClause(url *url.URL, expectedProtocol string) (string, error) {

	scheme := url.Scheme
	if expectedProtocol != scheme {
		return "", errors.New(
			`connection protocol (%s) value is not (%s)`,
			scheme,
			expectedProtocol,
		)
	}

	return scheme, nil
}
