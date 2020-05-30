package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/willabides/octo-go"
	"golang.org/x/oauth2"
)

func main() {
	err := os.MkdirAll("tmp", 0750)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	oauthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	))

	client := octo.NewClient(oauthClient)

	resp, err := client.ReposGetArchiveLink(ctx, &octo.ReposGetArchiveLinkReq{
		Owner:         "WillAbides",
		Repo:          "octo-go",
		ArchiveFormat: "tarball",
		Ref:           "master",
	})

	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create(filepath.Join("tmp", "octo-go.tar"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(output, resp.HTTPResponse().Body)
	if err != nil {
		log.Fatal(err)
	}

	err = resp.HTTPResponse().Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
