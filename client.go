package octo

import (
	"context"
	"fmt"
	"net/http"
)

// Client is a client for the GitHub API
type Client struct {
	opts []RequestOption
}

// NewClient returns a new Client
func NewClient(opt ...RequestOption) *Client {
	return &Client{
		opts: opt,
	}
}

func doRequest(ctx context.Context, requester httpRequester, opt ...RequestOption) (*response, error) {
	req, err := requester.httpRequest(ctx, opt...)
	if err != nil {
		return nil, err
	}
	ro, err := buildRequestOptions(opt)
	if err != nil {
		return nil, err
	}
	if ro.authProvider != nil {
		var authHeader string
		authHeader, err = ro.authProvider.AuthorizationHeader(ctx)
		if err != nil {
			return nil, fmt.Errorf("error setting authorization header: %v", err)
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

func (c *Client) doRequest(ctx context.Context, requester httpRequester, opt ...RequestOption) (*response, error) {
	opts := make([]RequestOption, len(c.opts), len(c.opts)+len(opt))
	copy(opts, c.opts)
	opts = append(opts, opt...)
	return doRequest(ctx, requester, opts...)
}

// SetRequestOptions sets options that will be used on all requests this client makes
func (c *Client) SetRequestOptions(opt ...RequestOption) {
	c.opts = append(c.opts, opt...)
}
