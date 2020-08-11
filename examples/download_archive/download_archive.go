package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/requests/repos"
)

func main() {
	err := os.MkdirAll("tmp", 0o750)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ghAuth := octo.WithPATAuth(os.Getenv("GITHUB_TOKEN"))

	resp, err := repos.DownloadTarballArchive(ctx, &repos.DownloadTarballArchiveReq{
		Owner: "WillAbides",
		Repo:  "octo-go",
		Ref:   "main",
	}, ghAuth)
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
