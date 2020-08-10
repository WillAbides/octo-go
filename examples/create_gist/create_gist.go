package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/options/auth"
	"github.com/willabides/octo-go/requests/gists"
)

func main() {
	ctx := context.Background()

	client := octo.NewClient(
		auth.WithPATAuth(os.Getenv("GITHUB_TOKEN")),
		octo.WithUserAgent("octo-go examples"),
	)

	createResp, err := gists.Create(ctx, &gists.CreateReq{
		RequestBody: gists.CreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]gists.CreateReqBodyFiles{
				"foo.md": {
					Content: octo.String(`# my header

my body
`),
				},
			},
		},
	}, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("don't forget to delete your new gist at %s\n", createResp.Data.HtmlUrl)
	fmt.Println("on second thought...I'll just delete it for you")

	deleteResp, err := gists.Delete(ctx, &gists.DeleteReq{
		GistId: createResp.Data.Id,
	}, client)
	if err != nil {
		log.Fatal(err)
	}

	if deleteResp.HTTPResponse().StatusCode != http.StatusNoContent {
		fmt.Println("something went wrong...you better delete it yourself.")
	}
}
