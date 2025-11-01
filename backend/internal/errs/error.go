package errs

import (
	"fmt"
	"net/http"
)

type CustomError interface {
	error
	Name() string
	Message() string
	Status() int
}

type baseError struct {
	name    string
	message string
	status  int
}

func (e *baseError) Error() string {
	return fmt.Sprintf("%s: %s", e.name, e.message)
}

func (e *baseError) Name() string    { return e.name }
func (e *baseError) Message() string { return e.message }
func (e *baseError) Status() int     { return e.status }

func newError(name, message string, status int) CustomError {
	return &baseError{name, message, status}
}

func NotFound(message string) CustomError {
	return newError(ErrNameNotFound, message, http.StatusNotFound)
}

func Unauthorized(message string) CustomError {
	return newError(ErrNameUnauthorized, message, http.StatusUnauthorized)
}

func Validation(message string) CustomError {
	return newError(ErrNameValidation, message, http.StatusBadRequest)
}

func Conflict(message string) CustomError {
	return newError(ErrNameConflict, message, http.StatusConflict)
}

func Internal(message string) CustomError {
	return newError(ErrNameInternal, message, http.StatusInternalServerError)
}
