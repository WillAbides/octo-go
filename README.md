# octo-go

[![godoc](https://godoc.org/github.com/WillAbides/octo-go?status.svg)](https://godoc.org/github.com/WillAbides/octo-go)
[![ci](https://github.com/WillAbides/octo-go/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/WillAbides/octo-go/actions?query=workflow%3Aci+branch%3Amaster+event%3Apush)

octo-go is an experimental client for GitHub's v3 API. It is generated from the openapi schema that GitHub covertly
 publishes at https://unpkg.com/browse/@github/openapi@latest/ (this schema is in alpha/prerelease state. if you want
 to use it yourself, do so at your own risk).
 
This is WIP. Don't depend on it.

Until I write more about it, you can get an idea of how it works in "./examples".

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

auth := octo.WithAppInstallationAuth(appID, installationID, key, &octo.AppsCreateInstallationTokenReqBody{
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
 these headers to enable paging. Every request has a `Rel()` method that will update the request to point to the
 relative link.
 
Let me demonstrate with a contrived example `GetMilestoneIssueTitles`. Stick around for a more concise version at
 afterward.

```go
// GetMilestoneIssueTitles returns the titles of all golang/go issues with a given milestone
func GetMilestoneIssueTitles(ctx context.Context, client octo.Client, milestone string) ([]string, error) {
	var result []string

	// a request to get all golang issues with the given milestone should have enough 
	// results to require paging
	req := &octo.IssuesListForRepoReq{
		Owner:     "golang",
		Repo:      "go",
		Milestone: &milestone,
	}
	
	// do the request and get the first page of results
	resp, err := client.IssuesListForRepo(ctx, req)
	if err != nil {
		return nil, err
	}

	// add the first page of titles to the result
	for _, issue := range *resp.Data {
		result = append(result, issue.Title)
	}

	// set req to point to the next page of results
	// ok will be false if there is no "next" link on the previous response
	ok := req.Rel(octo.RelNext, resp)
	
	// maybe there was just one page after all
	if !ok {
		return result, nil
	}
	
	// get the second page of results
	resp, err = client.IssuesListForRepo(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, issue := range *resp.Data {
		result = append(result, issue.Title)
	}
	
	// this is getting tiresome. let's handle the rest of the pages in a loop
	for req.Rel(octo.RelNext, resp) {
		resp, err = client.IssuesListForRepo(ctx, req)
		if err != nil {
			return nil, err
		}
		for _, issue := range *resp.Data {
			result = append(result, issue.Title)
		}
	}
	
	return result, nil
}
```

`GetMilestoneIssueTitles` was overly verbose for the sake of explanation. `GetMilestoneIssueTitles2` is more like
 what you should really be writing.

```go
// GetMilestoneIssueTitles2 does the same as GetMilestoneIssueTitles with a tighter loop
func GetMilestoneIssueTitles2(ctx context.Context, client octo.Client, milestone string) ([]string, error) {
	var result []string
	
	req := &octo.IssuesListForRepoReq{
		Owner:     "golang",
		Repo:      "go",
		Milestone: &milestone,
	}
	
	ok := true
	for ok {
		resp, err := client.IssuesListForRepo(ctx, req)
		if err != nil {
			return nil, err
		}
		for _, issue := range *resp.Data {
			result = append(result, issue.Title)
		}
		ok = req.Rel(octo.RelNext, resp)
	}
	return result, nil
}
```

In addition to `RelNext` to get the next page, you can also use `RelPrev` to get the next page. `RelFirst` and
 `RelLast` get the first and last page of results.

## Rate Limits

The GitHub API has a general rate limit of 5,000 requests per hour for most authenticated requests and 60 per hour per
 ip address for unauthenticated requests. More details are in the [API documentation](https://developer.github.com/v3/#rate-limiting).

To check your rate limit status, these methods are available on all octo-go responses (`resp`):

`resp.RateLimitRemaining()` - returns the number of requests remaining (or -1 if the header is missing)

`resp.RateLimitReset()` - returns the time when the rate limit will reset (or zero value if the header is missing)

`resp.RateLimit()` - returns the rate limit (or -1 if the header is missing)

You can also explicitly get your rate limit status with `octo.RateLimitGet()`

