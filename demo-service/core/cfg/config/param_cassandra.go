package config

import (
	"net/url"
	"strings"

	"github.com/gocql/gocql"

	"github.com/KWRI/demo-service/core/errors"
)

const (
	//
	// CassandraConnectionScheme cassandra schema.
	//
	CassandraConnectionScheme = "cassandra"

	//
	// CassandraDefaultConsistencyLevel default consistency level
	// when it is not specified in connection string via param "cl" (i.e. "&cl=Quorum")
	//
	CassandraDefaultConsistencyLevel = gocql.LocalQuorum
)

//
// Consistency is an alias for gocql.Consistency type.
//
type Consistency = gocql.Consistency

//
// CassandraConnectionInfoProvider declares all the connection info getters.
//
type CassandraConnectionInfoProvider interface {
	//
	// GetHosts returns the list of hosts.
	//
	GetHosts() []string

	//
	// GetKeyspace returns cassandraKeyspace value.
	//
	GetKeyspace() string

	//
	// GetDataCenter returns data center value.
	//
	GetDataCenter() string

	//
	// GetUser returns user value.
	//
	GetUser() string

	//
	// GetPassword returns cassandraPassword value.
	//
	GetPassword() string

	//
	// GetConnectionString returns a connection string as it is.
	//
	GetConnectionString() string

	//
	// IsAuthorizationRequired returns true if auth info is set in the connection string.
	//
	IsAuthorizationRequired() bool

	//
	// IsDCAware returns true if data center info was set.
	//
	IsDCAware() bool

	//
	// GetConsistencyLevel returns consistency level or "" if not set
	//
	GetConsistencyLevel() Consistency
}

//
// ParameterCassandraProvider is an interface for the cassandra parameter entry.
//
type ParameterCassandraProvider interface {
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
// CassandraInfoParam is a cassandra database connection info parameter.
//
type CassandraInfoParam struct {
	hosts       []string
	keyspace    string
	dataCenter  string
	user        string
	password    string
	consistency Consistency
	value       string

	baseParam
}

//
// NewCassandraInfoHostOnly returns a new instance of cassandra parameter object.
//
func NewCassandraInfoHostOnly(host string) *CassandraInfoParam {

	return &CassandraInfoParam{
		hosts: []string{host},
	}
}

//
// NewCassandraInfo returns a new instance of cassandra parameter object.
//
func NewCassandraInfo(name, usage string) *CassandraInfoParam {

	return &CassandraInfoParam{
		baseParam: baseParam{
			name:      name,
			usage:     usage,
			paramType: ParameterTypeCassandra,
		},
	}
}

//
// getValuePointer returns value link.
//
func (p *CassandraInfoParam) getValuePointer() *string {

	return &p.value
}

//
// GetHosts returns the list of hosts.
//
func (p *CassandraInfoParam) GetHosts() []string {

	return p.hosts
}

//
// GetKeyspace returns cassandra Keyspace value.
//
func (p *CassandraInfoParam) GetKeyspace() string {

	return p.keyspace
}

//
// GetDataCenter returns data center value.
//
func (p *CassandraInfoParam) GetDataCenter() string {

	return p.dataCenter
}

//
// GetUser returns user value.
//
func (p *CassandraInfoParam) GetUser() string {

	return p.user
}

//
// GetPassword returns cassandraPassword value.
//
func (p *CassandraInfoParam) GetPassword() string {

	return p.password
}

//
// IsDCAware returns true if data center info was set.
//
func (p *CassandraInfoParam) IsDCAware() bool {

	return "" != p.dataCenter
}

//
// IsAuthorizationRequired returns true if auth info is set in the connection string.
//
func (p *CassandraInfoParam) IsAuthorizationRequired() bool {

	return "" != p.user && "" != p.password
}

//
// GetConsistencyLevel returns consistency level.
//
func (p *CassandraInfoParam) GetConsistencyLevel() Consistency {

	return p.consistency
}

//
// GetConnectionString returns a connection string as it is.
//
func (p *CassandraInfoParam) GetConnectionString() string {

	return p.value
}

//
// validate validates the cassandra connection string parameter.
//
func (p *CassandraInfoParam) validate() error {

	connectionString := p.value

	if "" == connectionString {
		return ErrCassandraConnectionStringIsEmpty
	}

	connectionInfo, err := url.Parse(connectionString)
	if nil != err {
		return errors.Wrap(err, ErrCassandraConnectionStringIsIncorrect.WithMessage(
			"parse error for (%s)", connectionString,
		))
	}

	if _, err = validateURLScheme(connectionInfo, CassandraConnectionScheme); nil != err {
		return errors.Wrap(err, ErrCassandraSchemeIsIncorrect.WithMessage(
			"connection protocol (%s) is not (%s)", connectionInfo.Scheme, CassandraConnectionScheme,
		))
	}

	if p.hosts, err = validateCassandraHostsClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "cassandra hosts (%s) are invalid", connectionInfo.Host)
	}

	if p.keyspace, err = validateCassandraKeyspaceClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "cassandra keyspace (%s) is invalid", connectionInfo.Path)
	}

	if p.user, p.password, err = validateCassandraAuthorizationClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "cassandra authorization info is invalid")
	}

	p.dataCenter = validateCassandraDCClause(connectionInfo)

	if p.consistency, err = validateCassandraConsistencyLevelClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "cassandra consistency level is incorrect")
	}

	return nil
}

//
// validateCassandraHostsClause validates the hosts clause.
//
func validateCassandraHostsClause(url *url.URL) ([]string, error) {

	hosts := strings.Split(url.Host, ",")
	if 0 == len(hosts) {
		return nil, ErrCassandraHostsAreEmpty
	}

	return hosts, nil
}

//
// validateCassandraKeyspaceClause validates the cassandraKeyspace clause.
//
func validateCassandraKeyspaceClause(url *url.URL) (string, error) {

	keyspace := strings.Trim(url.Path, "/")
	if "" == keyspace {
		return "", ErrCassandraKeyspaceIsEmpty
	}

	return keyspace, nil
}

//
// validateCassandraAuthorizationClause validates the authorization clause.
//
func validateCassandraAuthorizationClause(url *url.URL) (string, string, error) {

	if url.User != nil {
		pwd, exists := url.User.Password()
		if !exists {
			return "", "", ErrCassandraPasswordIsEmpty
		}

		return url.User.Username(), pwd, nil
	}

	return "", "", nil
}

//
// validateCassandraDCClause validates the data-center clause.
//
func validateCassandraDCClause(url *url.URL) string {

	if "" != url.Query().Get("dc") {
		return url.Query().Get("dc")
	}

	return ""
}

//
// validateCassandraConsistencyLevelClause validates the cassandra consistency level clause,
// or returns default if not specified.
//
func validateCassandraConsistencyLevelClause(url *url.URL) (Consistency, error) {

	if "" != url.Query().Get("cl") {
		consistencyLevel, err := gocql.ParseConsistencyWrapper(url.Query().Get("cl"))
		if nil != err {
			return gocql.Any, ErrCassandraConsistencyLevelIsIncorrect
		}

		return consistencyLevel, nil
	}

	return CassandraDefaultConsistencyLevel, nil
}
