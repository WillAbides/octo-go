package octo

import (
	"context"
	"fmt"
	"net/http"
)

// Client is a client for the GitHub API
type Client struct {
	requestOpts requestOpts
}

// NewClient returns a new Client
func NewClient(opt ...RequestOption) (*Client, error) {
	ro, err := buildRequestOptions(opt)
	if err != nil {
		return nil, err
	}
	return &Client{
		requestOpts: ro,
	}, nil
}

func (c *Client) doRequest(ctx context.Context, requester httpRequester, opt ...RequestOption) (*response, error) {
	opts := make([]RequestOption, 0, len(opt)+1)
	opts = append(opts, resetOptions(c.requestOpts))
	opts = append(opts, opt...)
	req, err := requester.httpRequest(ctx, opts...)
	if err != nil {
		return nil, err
	}
	ro, err := buildRequestOptions(opts)
	if err != nil {
		return nil, err
	}
	if ro.authProvider != nil {
		var authHeader string
		authHeader, err = ro.authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, fmt.Errorf("error setting authorization header")
		}
		req.Header.Set("Authorization", authHeader)
	}
	httpClient := ro.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	httpResponse, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp := &response{
		httpResponse:  httpResponse,
		httpRequester: requester,
		opts:          ro,
	}

	err = errorCheck(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetRequestOptions sets options that will be used on all requests this client makes
func (c *Client) SetRequestOptions(opt ...RequestOption) error {
	for _, o := range opt {
		if o == nil {
			continue
		}
		err := o(&c.requestOpts)
		if err != nil {
			return err
		}
	}
	return nil
}
