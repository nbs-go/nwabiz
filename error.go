package nwabiz

import (
	"fmt"
)

const (
	UnhandledError = iota
	InvalidCredentialsError
	EmptyDataError
)

var errMessages = map[int]string{
	InvalidCredentialsError: "invalid credentials",
	EmptyDataError:          "returned list of data is returned empty",
}

func NewUnhandledError(err error) error {
	return &Error{
		Code:    UnhandledError,
		Message: err.Error(),
		Source:  err,
	}
}

func NewError(code int) error {
	msg, ok := errMessages[code]
	if !ok {
		return NewUnhandledError(fmt.Errorf("unhandled error. ErrorCode: %d", code))
	}

	err := Error{
		Code:    code,
		Message: msg,
	}

	return &err
}

type Error struct {
	Code    int
	Message string
	Source  error
}

func (e *Error) Error() string {
	return "nwabiz: " + e.Message
}
