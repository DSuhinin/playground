package config

//
// ParameterType configuration parameter type.
//
type ParameterType string

const (
	//
	// ParameterTypeString config parameter with "string" type.
	//
	ParameterTypeString ParameterType = "string"

	//
	// ParameterTypeBase64String config parameter with "base64" type.
	//
	ParameterTypeBase64String ParameterType = "base64string"

	//
	// ParameterTypeDuration config parameter with "duration" type.
	//
	ParameterTypeDuration ParameterType = "duration"

	//
	// ParameterTypeInteger config parameter with "integer" type.
	//
	ParameterTypeInteger ParameterType = "integer"

	//
	// ParameterTypeFloat64 config parameter with "float64" type.
	//
	ParameterTypeFloat64 ParameterType = "float64"

	//
	// ParameterTypeBool config parameter with "bool" type.
	//
	ParameterTypeBool ParameterType = "bool"

	//
	// ParameterTypeCassandra config parameter with "cassandra" type.
	//
	ParameterTypeCassandra ParameterType = "cassandra"

	//
	// ParameterTypeRedis config parameter with "redis" type.
	//
	ParameterTypeRedis ParameterType = "redis"

	//
	// ParameterTypeConsul config parameter with "consul" type.
	//
	ParameterTypeConsul ParameterType = "consul"

	//
	// ParameterTypeLogger config parameter with "logger" type.
	//
	ParameterTypeLogger ParameterType = "logger"
)
