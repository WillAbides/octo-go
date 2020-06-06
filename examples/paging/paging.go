package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/willabides/octo-go"
)

func main() {
	ctx := context.Background()

	req := &octo.IssuesListCommentsReq{
		Owner:       "golang",
		Repo:        "go",
		IssueNumber: 1,
		Since:       octo.ISOTimeString(time.Now().AddDate(-20, 0, 0)),
		PerPage:     octo.Int64(4),
	}

	ghAuth := octo.RequestPATAuth(os.Getenv("GITHUB_TOKEN"))

	fmt.Println("Comments from golang/go's first GitHub issue:")
	ok := true
	for ok {
		resp, err := octo.IssuesListComments(ctx, req, ghAuth)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Data != nil {
			for _, r := range *resp.Data {
				fmt.Printf("%s commented at %s: %s\n", r.User.Login, r.CreatedAt, r.HtmlUrl)
			}
		}
		ok = req.Rel(octo.RelNext, resp)
	}
}
