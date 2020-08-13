package errors

import (
	"net/http"
)

// RequestError is an error building an *http.Request
type RequestError struct {
	Message string
}

func (e *RequestError) Error() string {
	return e.Message
}

// ResponseError is an error from an *http.Response.
type ResponseError interface {
	HttpResponse() *http.Response
	Error() string
	Data() *ResponseErrorData // data from the error body if it can be unmarshalled
	IsClientError() bool      // true if the http status is in the 4xx range
	IsServerError() bool      // true if the http status is in the 5xx range
}

// ResponseErrorData all 4xx response bodies and maybe some 5xx should unmarshal to this
type ResponseErrorData struct {
	DocumentationUrl string `json:"documentation_url,omitempty"`
	Message          string `json:"message,omitempty"`
	Errors           []struct {
		Code     string `json:"code,omitempty"`
		Field    string `json:"field,omitempty"`
		Message  string `json:"message,omitempty"`
		Resource string `json:"resource,omitempty"`
	} `json:"errors,omitempty"`
}
