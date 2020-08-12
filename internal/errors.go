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

// ErrorCheck checks for error responses
func ErrorCheck(resp *http.Response, validStatuses []int) error {
	code := resp.StatusCode
	for _, wantStatus := range validStatuses {
		if code == wantStatus {
			return nil
		}
	}
	switch {
	case code >= 400 && code < 500:
		return newClientError(resp)
	case code >= 500 && code < 600:
		return newServerError(resp)
	}
	if isRedirectOnly(validStatuses) && code < 300 {
		return nil
	}
	msg := fmt.Sprintf("received unexpected http status code %d, expected codes are %v", code, validStatuses)
	return &errors.UnexpectedStatusCodeError{
		HTTPResponse: resp,
		Message:      msg,
	}
}

func newClientError(resp *http.Response) error {
	msg := fmt.Sprintf("client error %d", resp.StatusCode)
	errorData, err := unmarshalErrorData(resp)
	if err != nil {
		errorData = nil
	}
	if errorData != nil && errorData.Message != "" {
		msg += ": " + errorData.Message
	}
	return &errors.ClientError{
		HTTPResponse: resp,
		Message:      msg,
	}
}

func newServerError(resp *http.Response) error {
	msg := fmt.Sprintf("client error %d", resp.StatusCode)
	errorData, err := unmarshalErrorData(resp)
	if err != nil {
		errorData = nil
	}
	if errorData != nil && errorData.Message != "" {
		msg += ": " + errorData.Message
	}
	return &errors.ServerError{
		HTTPResponse: resp,
		Message:      msg,
	}
}

// ErrorData all 4xx response bodies and maybe some 5xx should unmarshal to this
type ErrorData struct {
	DocumentationUrl string           `json:"documentation_url,omitempty"`
	Message          string           `json:"message,omitempty"`
	Errors           []ErrorDataError `json:"errors,omitempty"`
}

// ErrorDataError an Error field in ErrorData
type ErrorDataError struct {
	Code     string `json:"code,omitempty"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message,omitempty"`
	Resource string `json:"resource,omitempty"`
}

func unmarshalErrorData(resp *http.Response) (*ErrorData, error) {
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
	var errorData ErrorData
	err := json.NewDecoder(bodyReader).Decode(&errorData)
	if err != nil {
		return nil, err
	}
	return &errorData, nil
}
