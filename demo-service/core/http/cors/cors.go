package cors

import (
	"net/http"
	"strings"

	"github.com/rs/cors"
)

//
// CORS header names.
//
const (
	headerAccessControlAllowedOrigins     = "Access-Control-Allow-Origin"
	headerAccessControlAllowedMethods     = "Access-Control-Allow-Methods"
	headerAccessControlAllowedCredentials = "Access-Control-Allow-Credentials"
	headerAccessControlAllowedHeaders     = "Access-Control-Allow-Headers"
	headerAccessControlExposedHeaders     = "Access-Control-Expose-Headers"
)

//
// WrapHTTPHandler wraps HTTP handler with CORS middleware to expose headers configured by opts.
//
func WrapHTTPHandler(h http.Handler, o OptionsGetter) http.Handler {

	c := cors.New(cors.Options{
		Debug:            o.GetCORSEnableDebug(),
		AllowCredentials: o.GetCORSAllowCredentials(),
		AllowedOrigins:   normalizeHeader(o.GetCORSAllowedOrigins()),
		AllowedMethods:   normalizeHeader(o.GetCORSAllowedMethods()),
		AllowedHeaders:   normalizeHeader(o.GetCORSAllowedHeaders()),
		ExposedHeaders:   normalizeHeader(o.GetCORSExposedHeaders()),
	})

	return c.Handler(h)
}

//
// normalizeHeader normalizes comma separated values to a list of strings
//
func normalizeHeader(s string) (headers []string) {
	hs := strings.Split(s, ",")
	for _, h := range hs {
		if trimmedHeader := strings.TrimSpace(h); "" != trimmedHeader {
			headers = append(headers, trimmedHeader)
		}
	}

	return headers
}
