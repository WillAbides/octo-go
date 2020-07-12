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
	err := os.MkdirAll("tmp", 0o750)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ghAuth := octo.WithPATAuth(os.Getenv("GITHUB_TOKEN"))

	resp, err := octo.ReposDownloadArchive(ctx, &octo.ReposDownloadArchiveReq{
		Owner:         "WillAbides",
		Repo:          "octo-go",
		ArchiveFormat: "tarball",
		Ref:           "master",
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
