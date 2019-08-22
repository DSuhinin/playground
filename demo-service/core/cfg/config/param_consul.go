package config

import (
	"net/url"

	"github.com/KWRI/demo-service/core/errors"
)

const (
	//
	// ConsulConnectionScheme consul schema.
	//
	ConsulConnectionScheme = "consul"
)

//
// ConsulConnectionInfoProvider declares all the connection info getters.
//
type ConsulConnectionInfoProvider interface {
	//
	// GetHost returns the Consul address in address:port format.
	//
	GetHost() string

	//
	// GetUser returns user value.
	//
	GetUser() string

	//
	// GetPassword returns password value.
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

	// IsDCAware returns true if data center info was set.
	//
	IsDCAware() bool

	//
	// GetDataCenter returns data center value.
	//
	GetDataCenter() string
}

//
// ParameterConsulProvider is an interface for the consul parameter entry.
//
type ParameterConsulProvider interface {
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
// ConsulInfoParam represents consul connection info parameter.
//
type ConsulInfoParam struct {
	host       string
	user       string
	password   string
	value      string
	dataCenter string

	baseParam
}

//
// NewConsulInfo returns a new instance of consul parameter object.
//
func NewConsulInfo(name, usage string) *ConsulInfoParam {

	return &ConsulInfoParam{
		baseParam: baseParam{
			name:      name,
			usage:     usage,
			paramType: ParameterTypeConsul,
		},
	}
}

//
// getValuePointer returns value link.
//
func (p *ConsulInfoParam) getValuePointer() *string {

	return &p.value
}

//
// GetHost returns the consul host.
//
func (p *ConsulInfoParam) GetHost() string {

	return p.host
}

//
// GetDataCenter returns data center value.
//
func (p *ConsulInfoParam) GetDataCenter() string {

	return p.dataCenter
}

//
// GetUser returns user value.
//
func (p *ConsulInfoParam) GetUser() string {

	return p.user
}

//
// GetPassword returns password value.
//
func (p *ConsulInfoParam) GetPassword() string {

	return p.password
}

//
// IsDCAware returns true if data center info was set.
//
func (p *ConsulInfoParam) IsDCAware() bool {

	return "" != p.dataCenter
}

//
// IsAuthorizationRequired returns true if auth info is set in the connection string.
//
func (p *ConsulInfoParam) IsAuthorizationRequired() bool {

	return "" != p.user && "" != p.password
}

//
// GetConnectionString returns a connection string as it is.
//
func (p *ConsulInfoParam) GetConnectionString() string {

	return p.value
}

//
// validate validates the consul connection string parameter.
//
func (p *ConsulInfoParam) validate() error {

	connectionString := p.value

	if "" == connectionString {
		return ErrConsulConnectionStringIsEmpty
	}

	connectionInfo, err := url.Parse(connectionString)
	if nil != err {
		return errors.Wrap(err, ErrConsulConnectionStringIsIncorrect.WithMessage(
			"parse error for (%s)", connectionString,
		))
	}

	if _, err = validateURLScheme(connectionInfo, ConsulConnectionScheme); nil != err {
		return errors.Wrap(err, ErrConsulSchemeIsIncorrect.WithMessage(
			"connection protocol (%s) is not (%s)", connectionInfo.Scheme, ConsulConnectionScheme,
		))
	}

	if p.host, err = validateConsulHostClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "consul host (%s) is invalid", connectionInfo.Host)
	}

	if p.user, p.password, err = validateConsulAuthorizationClause(connectionInfo); nil != err {
		return errors.WithMessage(err, "consul authorization info is invalid")
	}

	p.dataCenter = validateConsulDCClause(connectionInfo)

	return nil
}

//
// validateConsulHostClause validates the hosts clause.
//
func validateConsulHostClause(url *url.URL) (string, error) {

	if url.Host == "" {
		return "", ErrConsulHostIsEmpty
	}

	return url.Host, nil
}

// validateConsulAuthorizationClause validates the authorization clause.
//
func validateConsulAuthorizationClause(url *url.URL) (string, string, error) {

	if url.User != nil {
		pwd, exists := url.User.Password()
		if !exists {
			return "", "", ErrConsulPasswordIsEmpty
		}

		return url.User.Username(), pwd, nil
	}

	return "", "", nil
}

//
// validateConsulDCClause validates the data-center clause.
//
func validateConsulDCClause(url *url.URL) string {

	if "" != url.Query().Get("dc") {
		return url.Query().Get("dc")
	}

	return ""
}
