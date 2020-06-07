package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/willabides/octo-go"
)

func main() {
	ctx := context.Background()
	client := octo.NewClient(
		octo.WithPATAuth(os.Getenv("GITHUB_TOKEN")),
	)

	blockers, err := getReleaseBlockers(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if len(blockers) == 0 {
		fmt.Println("there are no open release blockers")
		return
	}
	fmt.Println("these are the open release blockers")
	for _, title := range blockers {
		fmt.Println(title)
	}
}

func getReleaseBlockers(ctx context.Context, client octo.Client) ([]string, error) {
	var result []string

	// Build the initial request.
	req := &octo.IssuesListForRepoReq{
		Owner:   "golang",
		Repo:    "go",
		Labels:  octo.String("release-blocker"),
	}

	// ok will be true as long as there is a next page.
	ok := true

	for ok {
		// Get a page of issues.
		resp, err := client.IssuesListForRepo(ctx, req)
		if err != nil {
			return nil, err
		}

		// Add issue titles to the result.
		for _, issue := range *resp.Data {
			result = append(result, issue.Title)
		}

		// Update req to point to the next page of results.
		// If there is no next page, req.Rel will return false and break the loop
		ok = req.Rel(octo.RelNext, resp)
	}
	return result, nil
}
