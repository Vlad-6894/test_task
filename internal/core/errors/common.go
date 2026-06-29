package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("NotFound")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
)
