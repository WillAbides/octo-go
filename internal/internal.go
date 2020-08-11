package internal

import (
	"context"
	"fmt"

	"github.com/willabides/octo-go/requests"
)

var operationAttributes map[string][]EndpointAttribute

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

// OperationHasAttribute returns true if the operation id the given attribute
func OperationHasAttribute(id string, attr EndpointAttribute) bool {
	attrs := operationAttributes[id]
	for _, a := range attrs {
		if attr == a {
			return true
		}
	}
	return false
}

// DoRequest performs an http request and returns a Response
func DoRequest(ctx context.Context, builder *RequestBuilder, opts *requests.Options) (*requests.Response, error) {
	req, err := builder.HTTPRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	// TODO: move this into builder.HTTPRequest
	authProvider := opts.AuthProvider()
	if authProvider != nil {
		var authHeader string
		authHeader, err = authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, fmt.Errorf("error setting authorization header: %v", err)
		}
		req.Header.Set("Authorization", authHeader)
	}
	httpResponse, err := opts.HttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	resp := requests.NewResponse(httpResponse)

	err = ErrorCheck(resp, builder)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// String returns a pointer to s
func String(s string) *string {
	return &s
}
