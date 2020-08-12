package errors

import (
	"net/http"
)

// UnexpectedStatusCodeError is returned when an unexpected status code is received, but
// the status is not in the 4xx or 5xx range.
type UnexpectedStatusCodeError struct {
	HTTPResponse *http.Response
	Message      string
}

func (e *UnexpectedStatusCodeError) Error() string {
	return e.Message
}

// ClientError is returned when the http status is in the 4xx range
type ClientError struct {
	HTTPResponse *http.Response
	Message      string
}

func (e *ClientError) Error() string {
	return e.Message
}

// ServerError is returned when the http status is in the 5xx range
type ServerError struct {
	HTTPResponse *http.Response
	Message      string
}

func (e *ServerError) Error() string {
	return e.Message
}
