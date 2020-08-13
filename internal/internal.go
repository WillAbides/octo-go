package internal

import (
	"net/http"
	"regexp"
)

// EndpointAttribute is an attribute for an endpoint
type EndpointAttribute int

var relLinkExp = regexp.MustCompile(`<(.+?)>\s*;\s*rel="([^"]*)"`)

// RelLink returns the content of lnk from the response's Link header or "" if it does not exist
func RelLink(resp *http.Response, lnk string) string {
	for _, link := range resp.Header.Values("Link") {
		for _, match := range relLinkExp.FindAllStringSubmatch(link, -1) {
			if match[2] == lnk {
				return match[1]
			}
		}
	}
	return ""
}

// String returns a pointer to s
func String(s string) *string {
	return &s
}

func isRedirectOnly(validStatuses []int) bool {
	if len(validStatuses) == 0 {
		return false
	}
	for _, vs := range validStatuses {
		if vs < 300 || vs > 399 {
			return false
		}
	}
	return true
}
