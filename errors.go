package goaccept

import (
	"fmt"
)

// This error occurs when no acceptable type can be found
type ErrorNotAcceptable struct {
	Found      []RequestedType
	Acceptable []string
}

// Serializes the error
func (e ErrorNotAcceptable) Error() string {
	return fmt.Sprintf("No acceptable content found; %v from request, only %v supported", e.Found, e.Acceptable)
}

// This error occurs when the accept value passed in can't be parsed
type ErrorBadAccept struct {
	Actual string
}

// Serializes the error
func (e ErrorBadAccept) Error() string {
	return fmt.Sprintf("Bad Accept header detected, unable to parse: '%s'", e.Actual)
}
