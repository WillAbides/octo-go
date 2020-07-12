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
