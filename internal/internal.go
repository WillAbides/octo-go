package internal

import (
	"net/http"
	"regexp"
)

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
