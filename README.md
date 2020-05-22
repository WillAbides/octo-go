# octo-go

[![godoc](https://godoc.org/github.com/WillAbides/octo-go?status.svg)](https://godoc.org/github.com/WillAbides/octo-go)
[![ci](https://github.com/WillAbides/octo-go/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/WillAbides/octo-go/actions?query=workflow%3Aci+branch%3Amaster+event%3Apush)

octo-go is an experimental client for GitHub's v3 API. It is generated from the opanapi schema that GitHub covertly
 publishes at https://unpkg.com/browse/@github/openapi@latest/. 
 
This is WIP. Don't depend on it.

Until I write more about it, you can get an idea of how it works in "./examples".

Here is the simple example:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/willabides/octo-go"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))

	req, err := octo.IssuesGetReq{
		Owner:               "golang",
		Repo:                "go",
		IssueNumber:         1,
	}.HTTPRequest(ctx)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("unexpected status code")
	}

	var issue octo.IssuesGetResponseBody200
	err = octo.UnmarshalResponseBody(resp, &issue)
	if err != nil {
		log.Fatal("unexpected status code")
	}

	fmt.Printf("golang/go's first issue is titled %q and has %d comments\n", issue.Title, issue.Comments)
}
```
