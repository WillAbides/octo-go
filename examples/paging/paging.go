package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/options"
	"github.com/willabides/octo-go/options/auth"
	"github.com/willabides/octo-go/requests/issues"
)

func main() {
	ctx := context.Background()
	client := octo.NewClient(
		auth.WithPATAuth(os.Getenv("GITHUB_TOKEN")),
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

func getReleaseBlockers(ctx context.Context, opts []options.Option) ([]string, error) {
	var result []string

	// Build the initial request.
	req := &issues.ListForRepoReq{
		Owner:  "golang",
		Repo:   "go",
		Labels: octo.String("release-blocker"),
	}

	// ok will be true as long as there is a next page.
	ok := true

	for ok {
		// Get a page of issues.
		resp, err := issues.ListForRepo(ctx, req, opts...)
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
