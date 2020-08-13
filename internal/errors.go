package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/willabides/octo-go/errors"
)

// NewResponseError returns a new *ResponseError
func NewResponseError(msg string, resp *http.Response) *ResponseError {
	data, err := unmarshalErrorData(resp)
	if err != nil {
		data = nil
	}
	return &ResponseError{
		resp: resp,
		msg:  msg,
		data: data,
	}
}

// ResponseError implements errors.ResponseError
type ResponseError struct {
	resp *http.Response
	msg  string
	data *errors.ResponseErrorData
}

// HttpResponse implements errors.ResponseError
func (r *ResponseError) HttpResponse() *http.Response {
	return r.resp
}

func (r *ResponseError) Error() string {
	msg := r.msg
	if r.data != nil && r.data.Message != "" {
		msg += ": " + r.data.Message
	}
	return msg
}

// Data implements errors.ResponseError
func (r *ResponseError) Data() *errors.ResponseErrorData {
	return r.data
}

// IsClientError implements errors.ResponseError
func (r *ResponseError) IsClientError() bool {
	return r.resp != nil && r.resp.StatusCode >= 400 && r.resp.StatusCode < 500
}

// IsServerError implements errors.ResponseError
func (r *ResponseError) IsServerError() bool {
	return r.resp != nil && r.resp.StatusCode >= 500 && r.resp.StatusCode < 600
}

// NewRequestError returns a new RequestError
func newRequestError(msg string) error {
	return &errors.RequestError{
		Message: msg,
	}
}

// ResponseErrorCheck checks for error responses
func ResponseErrorCheck(resp *http.Response, validStatuses []int) error {
	code := resp.StatusCode
	for _, wantStatus := range validStatuses {
		if code == wantStatus {
			return nil
		}
	}
	switch {
	case code >= 400 && code < 600:
		return NewResponseError(fmt.Sprintf("client error %d", resp.StatusCode), resp)
	case code >= 500 && code < 600:
		return NewResponseError(fmt.Sprintf("server error %d", resp.StatusCode), resp)
	}
	if isRedirectOnly(validStatuses) && code < 300 {
		return nil
	}
	msg := fmt.Sprintf("received unexpected http status code %d, expected codes are %v", code, validStatuses)
	return NewResponseError(msg, resp)
}

func unmarshalErrorData(resp *http.Response) (*errors.ResponseErrorData, error) {
	if resp.Body == nil {
		return nil, fmt.Errorf("no body")
	}
	var nextBody bytes.Buffer
	bodyReader := io.TeeReader(resp.Body, &nextBody)
	//nolint:errcheck // If there's an error draining the response body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = resp.Body.Close()
		resp.Body = ioutil.NopCloser(&nextBody)
	}()
	var errorData errors.ResponseErrorData
	err := json.NewDecoder(bodyReader).Decode(&errorData)
	if err != nil {
		return nil, err
	}
	return &errorData, nil
}
