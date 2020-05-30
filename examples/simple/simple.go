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

	oauthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))

	client := octo.NewClient(oauthClient)

	issue, err := client.IssuesGet(ctx, &octo.IssuesGetReq{
		Owner:       "golang",
		Repo:        "go",
		IssueNumber: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("golang/go's first issue is titled %q and has received %d comments\n", issue.Data.Title, issue.Data.Comments)
}
