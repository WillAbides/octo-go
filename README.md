# octo-go

[![godoc](https://godoc.org/github.com/WillAbides/octo-go?status.svg)](https://godoc.org/github.com/WillAbides/octo-go)
[![ci](https://github.com/WillAbides/octo-go/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/WillAbides/octo-go/actions?query=workflow%3Aci+branch%3Amain+event%3Apush)

octo-go is an experimental client for GitHub's v3 API. It is generated from the openapi schema published at 
https://github.com/github/rest-api-description

Project status: __BETA__

## Overview

For every API endpoint, octo-cli provides a request struct and a response struct. The request struct is used to build 
the http request, and the response struct is used to handle the api's response. You can use these structs as-is and 
handle all the http details yourself, or you can let octo-go do the request for you as well. Each endpoint also has a 
function that accepts the endpoints request struct and returns the response struct.

Let's use the `issues/create` endpoint as an example. You would use `issues.CreateReq` to build your request.

You can build a request like this:

```go
req := issues.CreateReq{
    Owner: "myorg",
    Repo:  "myrepo",
    RequestBody: issues.CreateReqBody{
        Title: octo.String("hello world"),
        Body:  octo.String("greetings from octo-cli"),
        Labels: []string{"test", "hello-world"},
    },
}
```

Then you can perform the request with:

```go
resp, err := issues.Create(ctx, &req)
```

And finally get the id of the newly created issue with:

```go
issueID := resp.Data.Id
```

## User Agent

GitHub requires all requests have a User-Agent header set. Octo-go sets it to `octo-go` by default, but please set it
 to the name of your program instead. Do that with the option `octo.WithUserAgent("my wonderful computer program")`.

## Authentication

In most situations, octo-go can handle the authentication, but you can also provide your own transport to set the
 Authentication header if you want.
 
### Personal Access Token

This is the simplest and most common way to authenticate.

```go
myToken := os.Getenv("GITHUB_TOKEN") // or however you want to provide your token

client := octo.NewClient(octo.WithPATAuth(myToken))
```

### GitHub App

If you want to authenticate as a GitHub App, octo can do that for you too. You need to provide the app's private key
 in PEM format along with your app's ID.

```go
appID := int64(1)
key, err := ioutil.ReadFile("appsecretkey.pem")
if err != nil {
    log.Fatal(err)
}
client := octo.NewClient(octo.WithAppAuth(appID, key))
```

### GitHub App Installation

To authenticate as a GitHub App Installation, you need the installation's ID along with the app's ID and private key.

```go
appID := int64(1)
installationID := int64(99)
key, err := ioutil.ReadFile("appsecretkey.pem")
if err != nil {
    log.Fatal(err)
}

instTokenClient := octo.NewClient(octo.WithAppAuth(appID, key))

auth := octo.WithAppInstallationAuth(installationID, instTokenClient, nil)
client := octo.NewClient(auth)
```

When authenticating as an App Installation, you can also limit the token's authorization to specific repositories and
 scopes by setting the request body used to create the token.
 
```go
appID := int64(1)
installationID := int64(99)
repoID := int64(12)
key, err := ioutil.ReadFile("appsecretkey.pem")
if err != nil {
    log.Fatal(err)
}

instTokenClient := octo.NewClient(octo.WithAppAuth(appID, key))

auth := octo.WithAppInstallationAuth(installationID, instTokenClient, &apps.CreateInstallationAccessTokenReqBody{
    Permissions: map[string]string{
        "deployments": "write",
        "content":     "read",
    },
    RepositoryIds: []int64{repoID},
})
client := octo.NewClient(auth)
```

## Pagination

The GitHub API supports paging through result sets using relative links in the Link header. Octo-go makes use of
 these headers to enable paging. Every response has the methods `RelLink(lnk string)` and `HasRelLink(lnk string)` 
 to get relative links. You can call this with `RelNext` for the next page of results, `RelPrev` for the previous
 page. `RelFirst` and `RelLast` point the first and last page of results.
 
Every request has a `Rel(lnk string, resp *ResponseType)` method that will update the request to point to a response's
 relative link.
 
Let me demonstrate with an example. `getReleaseBlockers` will page through all open golang/go issues that are labeled
 "release-blocker" and return their titles.

```go
func getReleaseBlockers(ctx context.Context, client octo.Client) ([]string, error) {
	var result []string

	// Build the initial request.
	req := &issues.ListForRepoReq{
		Owner:  "golang",
		Repo:   "go",
		Labels: octo.String("release-blocker"),
	}

	// ok will be true as long as there is a next page.
	for ok := true; ok; {
		// Get a page of issues.
		resp, err := client.Issues().ListForRepo(ctx, req)
		if err != nil {
			return nil, err
		}

		// Add issue titles to the result.
		for _, issue := range resp.Data {
			result = append(result, issue.Title)
		}

		// Update req to point to the next page of results.
		// If there is no next page, req.Rel will return false and break the loop
		ok = req.Rel(octo.RelNext, resp)
	}
	return result, nil
}
```

## Rate Limits

The GitHub API has a general rate limit of 5,000 requests per hour for most authenticated requests and 60 per hour per
 ip address for unauthenticated requests. More details are in the [API documentation](https://developer.github.com/v3/#rate-limiting).

To check your rate limit status, these methods are available on all octo-go responses (`resp`):

`resp.RateLimitRemaining()` - returns the number of requests remaining (or -1 if the header is missing)

`resp.RateLimitReset()` - returns the time when the rate limit will reset (or zero value if the header is missing)

`resp.RateLimit()` - returns the rate limit (or -1 if the header is missing)

You can also explicitly get your rate limit status with `octo.RateLimitGet()`

