package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/willabides/octo-go"
)

func main() {
	ctx := context.Background()

	client, err := octo.NewClient(
		octo.RequestPATAuth(os.Getenv("GITHUB_TOKEN")),
	)
	if err != nil {
		log.Fatal(err)
	}

	createResp, err := client.GistsCreate(ctx, &octo.GistsCreateReq{
		RequestBody: octo.GistsCreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]octo.GistsCreateReqBodyFiles{
				"foo.md": {
					Content: octo.String(`# my header

my body
`),
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("don't forget to delete your new gist at %s\n", createResp.Data.HtmlUrl)
	fmt.Println("on second thought...I'll just delete it for you")

	deleteResp, err := client.GistsDelete(ctx, &octo.GistsDeleteReq{
		GistId: createResp.Data.Id,
	})

	if err != nil {
		log.Fatal(err)
	}

	if deleteResp.HTTPResponse().StatusCode != http.StatusNoContent {
		fmt.Println("something went wrong...you better delete it yourself.")
	}

}
