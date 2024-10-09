package common

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrBadRequest   = errors.New("bad request")
)
