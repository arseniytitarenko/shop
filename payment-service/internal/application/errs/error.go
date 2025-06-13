package errs

import "errors"

var (
	InvalidRequest = errors.New("invalid request")
	InvalidID      = errors.New("invalid id")
)
