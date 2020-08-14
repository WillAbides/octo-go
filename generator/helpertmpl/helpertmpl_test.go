package helpertmpl

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go/requests"
)

func exOptions(o *requests.Options) []requests.Option {
	return []requests.Option{func(opts *requests.Options) {
		*opts = *o
	}}
}

func Test_requestHeaders(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := buildHTTPRequestOptions{}
		got := requestHeaders(b)
		want := http.Header{
			"Accept": []string{"application/vnd.github.v3+json"},
		}
		require.Equal(t, want, got)
	})

	t.Run("previews", func(t *testing.T) {
		b := buildHTTPRequestOptions{
			Previews: map[string]bool{
				"foo": true,
				"bar": true,
			},
		}
		got := requestHeaders(b)
		want := http.Header{
			"Accept": []string{"application/vnd.github.v3+json", "application/vnd.github.bar-preview+json", "application/vnd.github.foo-preview+json"},
		}
		require.Equal(t, want, got)
	})

	t.Run("required previews", func(t *testing.T) {
		opts := new(requests.Options)
		opts.SetRequiredPreviews(true)
		b := buildHTTPRequestOptions{
			RequiredPreviews: []string{"foo", "bar"},
			Previews: map[string]bool{
				"bar": true,
				"baz": false,
			},
			Options: exOptions(opts),
		}
		got := requestHeaders(b)
		want := http.Header{
			"Accept": []string{"application/vnd.github.v3+json", "application/vnd.github.bar-preview+json", "application/vnd.github.foo-preview+json"},
		}
		require.Equal(t, want, got)
	})

	t.Run("all previews", func(t *testing.T) {
		opts := new(requests.Options)
		opts.SetAllPreviews(true)
		b := buildHTTPRequestOptions{
			AllPreviews: []string{"foo", "bar"},
			Options:     exOptions(opts),
		}
		got := requestHeaders(b)
		want := http.Header{
			"Accept": []string{"application/vnd.github.v3+json", "application/vnd.github.bar-preview+json", "application/vnd.github.foo-preview+json"},
		}
		require.Equal(t, want, got)
	})

	t.Run("existing HeaderVals", func(t *testing.T) {
		b := buildHTTPRequestOptions{
			HeaderVals: map[string]*string{
				"aCcept": strPtr("application/json"),
				"baz":    strPtr("qux"),
				"nilly":  nil,
			},
			Previews: map[string]bool{
				"foo":   true,
				"bar":   true,
				"flood": false,
			},
		}
		got := requestHeaders(b)
		want := http.Header{
			"Accept": []string{"application/json", "application/vnd.github.bar-preview+json", "application/vnd.github.foo-preview+json"},
			"Baz":    []string{"qux"},
		}
		require.Equal(t, want, got)
	})
}

func Test_updateURLQuery(t *testing.T) {
	t.Run("overrites values", func(t *testing.T) {
		u, err := url.Parse("https://foo/bar?a=c&a=b&c=e")
		require.NoError(t, err)
		vals := url.Values{
			"b": []string{"foo", "bar"},
			"a": []string{"d"},
		}
		want := `https://foo/bar?a=d&b=foo&b=bar&c=e`
		updateURLQuery(u, vals)
		require.Equal(t, want, u.String())
	})

	t.Run("nil vals", func(t *testing.T) {
		u, err := url.Parse("https://foo/bar?a=c&a=b")
		require.NoError(t, err)
		want := "https://foo/bar?a=c&a=b"
		updateURLQuery(u, nil)
		require.Equal(t, want, u.String())
	})

	t.Run("no existing query", func(t *testing.T) {
		u, err := url.Parse("https://foo/bar")
		require.NoError(t, err)
		vals := url.Values{
			"b": []string{"foo", "bar"},
			"a": []string{"d"},
		}
		want := `https://foo/bar?a=d&b=foo&b=bar`
		updateURLQuery(u, vals)
		require.Equal(t, want, u.String())
	})
}
