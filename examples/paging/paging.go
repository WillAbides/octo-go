package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/willabides/octo-go"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))

	commentsReq := octo.IssuesListCommentsReq{
		Owner:       "golang",
		Repo:        "go",
		IssueNumber: 1,
		Since:       octo.ISOTimeString(time.Now().AddDate(-20, 0, 0)),
		PerPage:     octo.Int64(4),
	}

	req, err := commentsReq.HTTPRequest(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("these accounts commented on golang/go's first issue:")
	for {
		var resp *http.Response
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 200 {
			var result octo.IssuesListCommentsResponseBody200
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			for _, r := range result {
				fmt.Println(r.User.Login)
			}
		}
		req, err = octo.ResponseNextPageReq(ctx, req, resp)
		if err != nil {
			log.Fatal(err)
		}
		if req == nil {
			break
		}
	}
}
