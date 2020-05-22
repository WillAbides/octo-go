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
