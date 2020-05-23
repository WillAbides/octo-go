package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/willabides/octo-go"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))

	req, err := octo.GistsCreateReq{
		RequestBody: octo.GistsCreateReqBody{
			Description: octo.String("test gist, pls delete"),
			Public:      octo.Bool(false),
			Files: map[string]*octo.GistsCreateReqBodyFiles{
				"foo.md": {
					Content: octo.String(`# my header

my body
`),
				},
			},
		},
	}.HTTPRequest(ctx)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var gistInfo octo.GistsCreateResponseBody201
	err = octo.UnmarshalResponseBody(resp, &gistInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("don't forget to delete your new gist at %s\n", gistInfo.HtmlUrl)
	fmt.Println("on second thought...I'll just delete it for you")

	req, err = octo.GistsDeleteReq{
		GistId: gistInfo.Id,
	}.HTTPRequest(ctx)

	if err != nil {
		log.Fatal(err)
	}

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusNoContent {
		fmt.Println("something went wrong...you better delete it yourself.")
	}

}
