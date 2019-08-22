package config

//
// ParameterIntProvider is an interface for the Int parameter entry.
//
type ParameterFloat64Provider interface {
	//
	// ParameterProvider include base parameter interface.
	//
	ParameterProvider

	//
	// getValuePointer returns value link.
	//
	getValuePointer() *float64

	//
	// getDefault returns a string parameter value.
	//
	getDefault() float64

	//
	// getValue returns the value link for passing to the flags parsing block.
	//
	getValue() float64
}

//
// Float64Param is a float64 configuration parameter.
//
type Float64Param struct {
	value        float64
	defaultValue float64

	baseParam
}

//
// NewFloat64 returns a new instance of float64 parameter object.
//
func NewFloat64(name, usage string, defaultValue float64) *Float64Param {

	return &Float64Param{
		defaultValue: defaultValue,
		baseParam: baseParam{
			name:      name,
			usage:     usage,
			paramType: ParameterTypeFloat64,
		},
	}
}

//
// getValue returns an float64 value.
//
func (p *Float64Param) getValue() float64 {

	return p.value
}

//
// getDefault returns a default value.
//
func (p *Float64Param) getDefault() float64 {

	return p.defaultValue
}

//
// getValuePointer returns link to the float64 value.
//
func (p *Float64Param) getValuePointer() *float64 {

	return &p.value
}

//
// validate validates the float64 configuration parameter value not to be an empty string.
//
func (p *Float64Param) validate() error {

	return nil
}
