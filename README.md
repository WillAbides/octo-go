# octo-go

[![godoc](https://godoc.org/github.com/WillAbides/octo-go?status.svg)](https://godoc.org/github.com/WillAbides/octo-go)
[![ci](https://github.com/WillAbides/octo-go/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/WillAbides/octo-go/actions?query=workflow%3Aci+branch%3Amaster+event%3Apush)

octo-go is an experimental client for GitHub's v3 API. It is generated from the openapi schema that GitHub covertly
 publishes at https://unpkg.com/browse/@github/openapi@latest/ (this schema is in alpha/prerelease state. if you want
 to use it yourself, do so at your own risk).
 
This is WIP. Don't depend on it.

Until I write more about it, you can get an idea of how it works in "./examples".

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

## Rate Limits

The GitHub API has a general rate limit of 5,000 requests per hour for most authenticated requests and 60 per hour per
 ip address for unauthenticated requests. More details are in the API documention

To check your rate limit status, these methods are available on all octo-go responses (`resp`):

`resp.RateLimitRemaining()` - returns the number of requests remaining (or -1 if the header is missing)

`resp.RateLimitReset()` - returns the time when the rate limit will reset (or zero value if the header is missing)

`resp.RateLimit()` - returns the rate limit (or -1 if the header is missing)

You can also explicitly get your rate limit status with `octo.RateLimitGet()`

