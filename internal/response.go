package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/willabides/octo-go/common"
	"github.com/willabides/octo-go/options"
)

// DecodeResponseBody unmarshals a common body onto target
func DecodeResponseBody(r *common.Response, builder *RequestBuilder, opts *options.Options, target interface{}) error {
	if builder.HasAttribute(AttrRedirectOnly) {
		return nil
	}
	origBody := r.HTTPResponse().Body
	var bodyReader io.Reader = origBody
	if opts.PreserveResponseBody() {
		var buf bytes.Buffer
		bodyReader = io.TeeReader(r.HTTPResponse().Body, &buf)
		r.HTTPResponse().Body = ioutil.NopCloser(&buf)
	}
	//nolint:errcheck // If there's an error draining the common body, there was probably already an error reported.
	defer func() {
		_, _ = ioutil.ReadAll(bodyReader)
		_ = origBody.Close()
	}()
	if !statusCodeInList(r, builder.DataStatuses) {
		return nil
	}
	if target == nil {
		return nil
	}
	return json.NewDecoder(bodyReader).Decode(target)
}

func statusCodeInList(r *common.Response, codes []int) bool {
	if r.HTTPResponse() == nil {
		return false
	}
	for _, code := range codes {
		if r.HTTPResponse().StatusCode == code {
			return true
		}
	}
	return false
}

// SetBoolResult sets the value of ptr to true if r has a 204 status code to true or false if the status code is 404
//  returns an error if the common is any other value
func SetBoolResult(r *common.Response, ptr *bool) error {
	switch r.HTTPResponse().StatusCode {
	case 204:
		*ptr = true
	case 404:
		*ptr = false
	default:
		return fmt.Errorf("non-boolean common status")
	}
	return nil
}
