package requests

// RequestError is an error building an *http.Request
type RequestError struct {
	Message string
}

func (e *RequestError) Error() string {
	return e.Message
}
