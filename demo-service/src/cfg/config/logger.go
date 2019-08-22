package config

//
// GetLoggerLevel returns defined logging level.
//
func (p *Config) GetLoggerLevel() string {
	return p.GetLogLevel(ConfLogLevel)
}
