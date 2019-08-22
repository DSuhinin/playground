package config

import (
	"net/url"

	"github.com/KWRI/demo-service/core/errors"
)

//
// validateURLScheme validates the connection protocol clause.
//
func validateURLScheme(url *url.URL, expectedScheme string) (string, error) {
	scheme := url.Scheme
	if expectedScheme != scheme {
		return "", errors.New(
			`connection protocol (%s) value is not (%s)`,
			scheme,
			expectedScheme,
		)
	}

	return scheme, nil
}
