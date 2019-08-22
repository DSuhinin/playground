package cors

//
// OptionsGetter is an interface that provides CORS Access Control header values.
//
type OptionsGetter interface {
	//
	// GetCORSEnableDebug returns a Enable/Disable status for CORS debug.
	//
	GetCORSEnableDebug() bool

	//
	// GetCORSAllowedOrigins returns a CORS Allow Origin header values separated by comma.
	//
	GetCORSAllowedOrigins() string

	//
	// GetCORSAllowedMethods returns a CORS Allow Methods header values separated by comma.
	//
	GetCORSAllowedMethods() string

	//
	// GetCORSAllowedHeaders returns a CORS Allow Headers header values separated by comma.
	//
	GetCORSAllowedHeaders() string

	//
	// GetCORSExposedHeaders returns a CORS Expose Headers header values separated by comma.
	//
	GetCORSExposedHeaders() string

	//
	// GetCORSAllowCredentials returns a CORS Allow Credentials header value.
	//
	GetCORSAllowCredentials() bool
}
