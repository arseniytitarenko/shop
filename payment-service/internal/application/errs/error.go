package errs

import "errors"

var (
	InvalidRequest          = errors.New("invalid request")
	ErrAccountAlreadyExists = errors.New("account already exists")
)
