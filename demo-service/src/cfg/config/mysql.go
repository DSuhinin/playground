package config

//
// GetMySQLDSN returns MySQL DSN connection string.
//
func (p *Config) GetMySQLDSN() string {
	return p.GetString(ConfMySQLDSN)
}
