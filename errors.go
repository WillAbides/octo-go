package octo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/willabides/octo-go/internal"
)

func errorCheck(resp *response) error {
	for _, fn := range []func(*response) error{
		clientErrorCheck,
		serverErrorCheck,
		unexpectedStatusCheck,
	} {
		err := fn(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnexpectedStatusCodeError is returned when an unexpected status code is received, but
// the status is not in the 4xx or 5xx range.
type UnexpectedStatusCodeError struct {
	wantedCodes []int
	gotCode     int
	response
}

func (e *UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("received unexpected http status code %d, expected codes are %v", e.gotCode, e.wantedCodes)
}

func unexpectedStatusCheck(resp *response) error {
	valid := resp.httpRequester.validStatuses()
	opID := internal.ReqOperationID(resp.httpRequester)
	if internal.OperationHasAttribute(opID, internal.AttrRedirectOnly) {
		return nil
	}
	if internal.OperationHasAttribute(opID, internal.AttrBoolean) {
		valid = append(valid, 404)
	}
	statusCode := resp.httpResponse.StatusCode
	if resp.statusCodeInList(valid) {
		return nil
	}
	return &UnexpectedStatusCodeError{
		wantedCodes: valid,
		gotCode:     statusCode,
		response:    *resp,
	}
}

// ClientError is returned when the http status is in the 4xx range
type ClientError struct {
	response
	ErrorData *ErrorData
}

func (e *ClientError) Error() string {
	if e.ErrorData == nil || e.ErrorData.Message == "" {
		return fmt.Sprintf("client error %d", e.response.httpResponse.StatusCode)
	}
	return fmt.Sprintf("client error %d: %s", e.response.httpResponse.StatusCode, e.ErrorData.Message)
}

func clientErrorCheck(resp *response) error {
	statusCode := resp.httpResponse.StatusCode
	if statusCode < 400 || statusCode > 499 {
		return nil
	}
	opID := internal.ReqOperationID(resp.httpRequester)

	// 404 isn't an error for boolean endpoints ¯\_(ツ)_/¯
	if internal.OperationHasAttribute(opID, internal.AttrBoolean) && statusCode == 404 {
		return nil
	}

	clientErr := &ClientError{
		response:  *resp,
		ErrorData: new(ErrorData),
	}
	err := clientErr.ErrorData.decode(resp.httpResponse)
	if err != nil {
		clientErr.ErrorData = nil
	}
	return clientErr
}

// ServerError is returned when the http status is in the 5xx range
type ServerError struct {
	response
	ErrorData *ErrorData
}

func (e *ServerError) Error() string {
	if e.ErrorData == nil || e.ErrorData.Message == "" {
		return fmt.Sprintf("client error %d", e.response.httpResponse.StatusCode)
	}
	return fmt.Sprintf("client error %d: %s", e.response.httpResponse.StatusCode, e.ErrorData.Message)
}

func serverErrorCheck(resp *response) error {
	statusCode := resp.httpResponse.StatusCode
	if statusCode < 500 || statusCode > 599 {
		return nil
	}
	serverErr := &ServerError{
		response:  *resp,
		ErrorData: new(ErrorData),
	}
	err := serverErr.ErrorData.decode(resp.httpResponse)
	if err != nil {
		serverErr.ErrorData = nil
	}
	return serverErr
}

// ErrorData all 4xx response bodies and maybe some 5xx should unmarshal to this
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
	//nolint:errcheck // If there's an error draining the response body, there was probably already an error reported.
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
