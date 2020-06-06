# octo-go

[![godoc](https://godoc.org/github.com/WillAbides/octo-go?status.svg)](https://godoc.org/github.com/WillAbides/octo-go)
[![ci](https://github.com/WillAbides/octo-go/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/WillAbides/octo-go/actions?query=workflow%3Aci+branch%3Amaster+event%3Apush)

octo-go is an experimental client for GitHub's v3 API. It is generated from the openapi schema that GitHub covertly
 publishes at https://unpkg.com/browse/@github/openapi@latest/ (this schema is in alpha/prerelease state. if you want
 to use it yourself, do so at your own risk).
 
This is WIP. Don't depend on it.

Until I write more about it, you can get an idea of how it works in "./examples".

Here is the simple example:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/willabides/octo-go"
)

func main() {
	ctx := context.Background()

	issue, err := octo.IssuesGet(ctx, &octo.IssuesGetReq{
		Owner:       "golang",
		Repo:        "go",
		IssueNumber: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("golang/go's first issue is titled %q and has received %d comments\n", issue.Data.Title, issue.Data.Comments)
}
```
