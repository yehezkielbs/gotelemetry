package gotelemetry

import (
	"fmt"
)

type Error struct {
	StatusCode int         // HTTP status code of the error
	Message    string      // Error message
	Data       interface{} // Additional error data returned by the server
}

func NewError(statusCode int, message string) *Error {
	return &Error{statusCode, message, nil}
}

func NewErrorWithData(statusCode int, message string, data interface{}) *Error {
	return &Error{statusCode, message, data}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d -> %s\n%s", e.StatusCode, e.Message, e.Data)
}
