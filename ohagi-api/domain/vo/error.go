package vo

import "errors"

var (
	ErrEnum      = errors.New("enum error")
	ErrMaxLength = errors.New("max length error")
	ErrEmpty     = errors.New("empty error")
)
