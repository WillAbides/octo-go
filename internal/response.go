package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DecodeResponseBody unmarshalls a response body onto target
func DecodeResponseBody(r *http.Response, target interface{}, preserveResponseBody bool) error {
	body := r.Body
	bb, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading body")
	}
	err = body.Close()
	if err != nil {
		return fmt.Errorf("error closing body")
	}
	if preserveResponseBody {
		r.Body = ioutil.NopCloser(bytes.NewReader(bb))
	}
	return json.Unmarshal(bb, &target)
}

// IntInSlice returns true if i is in want
func IntInSlice(i int, want []int) bool {
	for _, code := range want {
		if i == code {
			return true
		}
	}
	return false
}

// SetBoolResult sets the value of ptr to true if r has a 204 status code to true or false if the status code is 404
//  returns an error if the response is any other value
func SetBoolResult(r *http.Response, ptr *bool) error {
	switch r.StatusCode {
	case 204:
		*ptr = true
	case 404:
		*ptr = false
	default:
		return fmt.Errorf("non-boolean response status")
	}
	return nil
}
