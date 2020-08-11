package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/willabides/octo-go/requests"
)

// ErrorCheck checks for errors in the common
func ErrorCheck(resp *requests.Response, builder *RequestBuilder) error {
	err := clientErrorCheck(resp, builder)
	if err != nil {
		return err
	}
	err = serverErrorCheck(resp)
	if err != nil {
		return err
	}
	err = unexpectedStatusCheck(resp, builder)
	if err != nil {
		return err
	}
	return nil
}

// UnexpectedStatusCodeError is returned when an unexpected status code is received, but
// the status is not in the 4xx or 5xx range.
type UnexpectedStatusCodeError struct {
	wantedCodes []int
	gotCode     int
	requests.Response
}

func (e *UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("received unexpected http status code %d, expected codes are %v", e.gotCode, e.wantedCodes)
}

func unexpectedStatusCheck(resp *requests.Response, builder *RequestBuilder) error {
	valid := make([]int, len(builder.ValidStatuses))
	copy(valid, builder.ValidStatuses)
	if builder.HasAttribute(AttrBoolean) {
		valid = append(valid, 404)
	}
	if builder.HasAttribute(AttrRedirectOnly) {
		return nil
	}
	statusCode := resp.HTTPResponse().StatusCode
	if statusCodeInList(resp, valid) {
		return nil
	}
	return &UnexpectedStatusCodeError{
		wantedCodes: valid,
		gotCode:     statusCode,
		Response:    *resp,
	}
}

// ClientError is returned when the http status is in the 4xx range
type ClientError struct {
	requests.Response
	ErrorData *ErrorData
}

func (e *ClientError) Error() string {
	if e.ErrorData == nil || e.ErrorData.Message == "" {
		return fmt.Sprintf("client error %d", e.Response.HTTPResponse().StatusCode)
	}
	return fmt.Sprintf("client error %d: %s", e.Response.HTTPResponse().StatusCode, e.ErrorData.Message)
}

func clientErrorCheck(resp *requests.Response, builder *RequestBuilder) error {
	statusCode := resp.HTTPResponse().StatusCode
	if statusCode < 400 || statusCode > 499 {
		return nil
	}

	// 404 isn't an error for boolean endpoints ¯\_(ツ)_/¯
	if builder.HasAttribute(AttrBoolean) && statusCode == 404 {
		return nil
	}

	clientErr := &ClientError{
		Response:  *resp,
		ErrorData: new(ErrorData),
	}
	err := clientErr.ErrorData.decode(resp.HTTPResponse())
	if err != nil {
		clientErr.ErrorData = nil
	}
	return clientErr
}

// ServerError is returned when the http status is in the 5xx range
type ServerError struct {
	requests.Response
	ErrorData *ErrorData
}

func (e *ServerError) Error() string {
	if e.ErrorData == nil || e.ErrorData.Message == "" {
		return fmt.Sprintf("client error %d", e.Response.HTTPResponse().StatusCode)
	}
	return fmt.Sprintf("client error %d: %s", e.Response.HTTPResponse().StatusCode, e.ErrorData.Message)
}

func serverErrorCheck(resp *requests.Response) error {
	statusCode := resp.HTTPResponse().StatusCode
	if statusCode < 500 || statusCode > 599 {
		return nil
	}
	serverErr := &ServerError{
		Response:  *resp,
		ErrorData: new(ErrorData),
	}
	err := serverErr.ErrorData.decode(resp.HTTPResponse())
	if err != nil {
		serverErr.ErrorData = nil
	}
	return serverErr
}

// ErrorData all 4xx common bodies and maybe some 5xx should unmarshal to this
type ErrorData struct {
	DocumentationUrl string           `json:"documentation_url,omitempty"`
	Message          string           `json:"message,omitempty"`
	Errors           []ErrorDataError `json:"errors,omitempty"`
}

func (e *ErrorData) decode(resp *http.Response) error {
	if resp.Body == nil {
		return fmt.Errorf("no body")
	}
	var nextBody bytes.Buffer
	bodyReader := io.TeeReader(resp.Body, &nextBody)
	//nolint:errcheck // If there's an error draining the common body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = resp.Body.Close()
		resp.Body = ioutil.NopCloser(&nextBody)
	}()
	return json.NewDecoder(bodyReader).Decode(e)
}

// ErrorDataError an Error field in ErrorData
type ErrorDataError struct {
	Code     string `json:"code,omitempty"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message,omitempty"`
	Resource string `json:"resource,omitempty"`
}
