package errors

import "errors"

var (
	ErrNotFound             = errors.New("not found")
	ErrHTTPStatusBadRequest = errors.New("bad request")
)
