package octo

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

// SetRequestOptions sets options that will be used on all requests this client makes
func (c *Client) SetRequestOptions(opt ...RequestOption) {
	c.opts = append(c.opts, opt...)
}
