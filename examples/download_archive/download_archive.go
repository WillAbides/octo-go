package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/willabides/octo-go"
)

func main() {
	err := os.MkdirAll("tmp", 0750)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := octo.NewClient(
		octo.RequestPATAuth(os.Getenv("GITHUB_TOKEN")),
	)
	if err != nil {
		log.Fatal(err)
	}

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
