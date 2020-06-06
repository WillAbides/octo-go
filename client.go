package octo

// Client is a client for the GitHub API
type Client []RequestOption

// NewClient returns a new Client
func NewClient(opt ...RequestOption) Client {
	return Client(opt)
}
