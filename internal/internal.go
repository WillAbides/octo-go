package internal

import (
	"context"

	"github.com/willabides/octo-go/requests"
)

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

// DoRequest performs an http request and returns a Response
func DoRequest(ctx context.Context, builder *RequestBuilder, opts *requests.Options) (*requests.Response, error) {
	req, err := builder.HTTPRequest(ctx, opts)
	if err != nil {
		return nil, err
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
