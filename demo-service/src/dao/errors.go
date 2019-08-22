package dao

import "github.com/KWRI/demo-service/core/errors"

//
// Database errors.
// nolint
//
var (
	ErrorInternalError  = errors.New("database internal error")
	ErrorEntityNotFound = errors.New("sql: no rows in result set")
)
