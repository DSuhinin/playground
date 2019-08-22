package app

import (
	"github.com/KWRI/demo-service/core/errors"
)

//
// Service errors.
// nolint
//
var (
	// HTTP 500 errors
	ErrInternalError = errors.NewHTTP500Error(
		10000,
		"internal error happened.",
	)

	// Service errors.
	ErrInvalidJson = errors.NewHTTP400Error(
		40000,
		"invalid json",
	)

	ErrDealNotFound = errors.NewHTTP404Error(
		40001,
		"deal not found.",
	)
)
