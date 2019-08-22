package cors

import (
	"net/http"
)

//
// handlerStub is a stub HTTP handlerStub.
//
type handlerStub struct {
}

//
// ServeHTTP is a stub HTTP serving method.
//
func (h *handlerStub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

//
// optsGetterStub is a CORS options getter stub
//
type optsGetterStub struct {
	enableDebug      bool
	allowedOrigins   string
	allowedMethods   string
	allowedHeaders   string
	exposedHeaders   string
	allowCredentials bool
}

//
// GetCORSEnableDebug method stub.
//
func (o *optsGetterStub) GetCORSEnableDebug() bool {

	return o.enableDebug
}

//
// GetCORSAllowedOrigins method stub.
//
func (o *optsGetterStub) GetCORSAllowedOrigins() string {

	return o.allowedOrigins
}

//
// GetCORSAllowedMethods method stub.
//
func (o *optsGetterStub) GetCORSAllowedMethods() string {

	return o.allowedMethods
}

//
// GetCORSAllowedHeaders method stub.
//
func (o *optsGetterStub) GetCORSAllowedHeaders() string {

	return o.allowedHeaders
}

//
// GetCORSExposedHeaders method stub.
//
func (o *optsGetterStub) GetCORSExposedHeaders() string {

	return o.exposedHeaders
}

//
// GetCORSAllowCredentials method stub.
//
func (o *optsGetterStub) GetCORSAllowCredentials() bool {

	return o.allowCredentials
}
