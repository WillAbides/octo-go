package internal

import (
	"context"
	"fmt"
	"reflect"

	"github.com/willabides/octo-go/options"
)

var (
	reqOperationIDs     map[string]string //nolint:unused // save for later
	operationAttributes map[string][]EndpointAttribute
)

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

// OperationAttributes returns the EndpointAttributes associated with an operation id
func OperationAttributes(id string) []EndpointAttribute {
	return operationAttributes[id]
}

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

// structName returns the name of a struct from its reflect type or a pointer
//nolint:unused // save for later
func structName(tp reflect.Type) string {
	if tp.Kind() == reflect.Ptr {
		return structName(tp.Elem())
	}
	return tp.Name()
}

// DoRequest performs an http request and returns a Response
func DoRequest(ctx context.Context, builder *RequestBuilder, opts *options.Options) (*Response, error) {
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
	resp := &Response{
		opts:         opts,
		httpResponse: httpResponse,
		reqBuilder:   builder,
	}

	err = ErrorCheck(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
