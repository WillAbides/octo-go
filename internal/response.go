package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// UnmarshalResponseBody unmarshalls a response body onto target. Non-nil errors will have the type *errors.ResponseError.
func UnmarshalResponseBody(r *http.Response, target interface{}) error {
	body := r.Body
	bb, err := ioutil.ReadAll(body)
	if err != nil {
		return NewResponseError("could not read response body", r)
	}
	err = body.Close()
	if err != nil {
		return NewResponseError("could not close response body", r)
	}
	err = json.Unmarshal(bb, &target)
	if err != nil {
		return NewResponseError("could not unmarshal json from response body", r)
	}
	return nil
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
		return NewResponseError("non-boolean response status", r)
	}
	return nil
}
