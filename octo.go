package octo

import (
	"time"

	"github.com/willabides/octo-go/requests"
)

//go:generate go run ./generator -schema "api.github.com.json" -pkgpath "github.com/willabides/octo-go" -pkg octo

// Common values for rel links
const (
	RelNext  = "next"
	RelPrev  = "prev"
	RelFirst = "first"
	RelLast  = "last"
)

// ISOTimeString returns a pointer to tm formated as an iso8601/rfc3339 string
func ISOTimeString(tm time.Time) *string {
	return String(tm.Format(time.RFC3339))
}

// String returns a pointer to s
func String(s string) *string {
	return &s
}

// Bool returns a pointer to b
func Bool(b bool) *bool {
	return &b
}

// Int64 returns a pointer to i
func Int64(i int64) *int64 {
	return &i
}

// Client is a set of options to apply to requests
type Client []requests.Option

// NewClient returns a new Client
func NewClient(opt ...requests.Option) Client {
	return opt
}
