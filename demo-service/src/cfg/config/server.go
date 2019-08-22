package config

import "time"

//
// GetServerHTTPAddress returns service HTTP address.
//
func (p *Config) GetServerHTTPAddress() string {
	return p.GetString(ConfServerAddress)
}

//
// GetServerReadTimeout returns an HTTP Server Read timeout.
//
func (p *Config) GetServerReadTimeout() time.Duration {
	return p.GetDuration(ConfServerReadTimeout)
}

//
// GetServerWriteTimeout returns an HTTP Server Write timeout.
//
func (p *Config) GetServerWriteTimeout() time.Duration {
	return p.GetDuration(ConfServerWriteTimeout)
}
