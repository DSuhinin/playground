package helper

import (
	"github.com/KWRI/demo-service/core/errors"
)

//
// ExtractHTTPError is a utility function to check is any of errors inside the error stack passed as an err parameter
// contains the error of type expectedError.
//
func ExtractHTTPError(err error) errors.HTTPError {

	httpError := errors.Cause(err, (*errors.HTTPError)(nil))
	if nil == httpError {
		return errors.HTTPError{}
	}

	return httpError.(errors.HTTPError)
}
