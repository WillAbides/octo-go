package octo

import (
	"context"
)

// AuthProvider sets the Authorization header authenticate you with the API
type AuthProvider interface {
	AuthorizationHeader(ctx context.Context) (string, error)
}
