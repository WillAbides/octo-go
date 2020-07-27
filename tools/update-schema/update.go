package main

import (
	"archive/zip"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/components"
)

func main() {
	ctx := context.Background()

	var outputPath string
	var workingDir string
	var releaseTag string
	var force bool

	flag.StringVar(&outputPath, "out", "api.github.com.json", "where to write the schema")
	flag.StringVar(&workingDir, "workdir", ".", "directory of operations")
	flag.StringVar(&releaseTag, "tag", "latest", "tag to fetch")
	flag.BoolVar(&force, "force", false, "force download even if we already have the version")

	flag.Parse()

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN is required")
	}
	client := octo.NewClient(octo.WithPATAuth(githubToken))

	tag := releaseTag
	var err error
	var rel *components.Release
	if tag == "latest" {
		rel, err = latestRelease(ctx, client)
		if err != nil {
			log.Fatal(err)
		}
		tag = rel.TagName
	}

	currentVersion, err := getCurrentVersion(workingDir)
	if err != nil {
		log.Fatal(err)
	}
	if currentVersion == tag && !force {
		fmt.Println("already up to date")
		return
	}

	if rel == nil {
		rel, err = taggedRelease(ctx, client, tag)
		if err != nil {
			log.Fatal(err)
		}
	}

	var asset *components.ReleaseAssetsItem
	assets := rel.Assets
	for i := range assets {
		if assets[i].Name == "descriptions.zip" {
			asset = &assets[i]
			break
		}
	}
	if asset == nil {
		log.Fatal("no descriptions.zip in release")
	}
	dlResp, err := http.Get(asset.BrowserDownloadUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = handleDownloadResponse(dlResp, "descriptions/api.github.com/api.github.com.json", filepath.Join(workingDir, "api.github.com.json"))
	if err != nil {
		log.Fatal(err)
	}
	err = setCurrentVersion(workingDir, tag)
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentVersion(workingDir string) (string, error) {
	versionFile := filepath.Join(workingDir, "current-schema-version.txt")
	b, err := ioutil.ReadFile(versionFile)
	switch {
	case err == nil:
		return string(b), nil
	case os.IsNotExist(err):
		return "v0.0.0", nil
	default:
		return "", err
	}
}

func setCurrentVersion(workingDir, version string) error {
	versionFile := filepath.Join(workingDir, "current-schema-version.txt")
	return ioutil.WriteFile(versionFile, []byte(version), 0o640) //nolint:gosec // 640 is fine
}

func latestRelease(ctx context.Context, client octo.Client) (*components.Release, error) {
	resp, err := client.ReposGetLatestRelease(ctx, &octo.ReposGetLatestReleaseReq{
		Owner: "github",
		Repo:  "rest-api-description",
	})
	if err != nil {
		return nil, err
	}
	rel := components.Release(resp.Data)
	return &rel, nil
}

func taggedRelease(ctx context.Context, client octo.Client, tag string) (*components.Release, error) {
	resp, err := client.ReposGetReleaseByTag(ctx, &octo.ReposGetReleaseByTagReq{
		Owner: "github",
		Repo:  "rest-api-description",
		Tag:   tag,
	})
	if err != nil {
		return nil, err
	}
	rel := components.Release(resp.Data)
	return &rel, nil
}

func handleDownloadResponse(resp *http.Response, zipPath, outputPath string) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non 200 reponse")
	}
	defer func() {
		_ = resp.Body.Close() //nolint:errcheck // no worries
	}()
	var bodyBuf bytes.Buffer
	_, err := io.Copy(&bodyBuf, resp.Body)
	if err != nil {
		return err
	}
	bodyRdr := bytes.NewReader(bodyBuf.Bytes())
	zipRdr, err := zip.NewReader(bodyRdr, bodyRdr.Size())
	if err != nil {
		return err
	}
	var fileRdr io.ReadCloser
	var unzipSize int64
	for _, file := range zipRdr.File {
		if file.Name != zipPath {
			continue
		}
		unzipSize = int64(file.UncompressedSize64)
		fileRdr, err = file.Open()
		if err != nil {
			return err
		}
		break
	}
	if fileRdr == nil {
		return fmt.Errorf("zip doesn't contain %s", zipPath)
	}
	outfile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = outfile.Close() //nolint:errcheck // no worries
	}()
	_, err = io.CopyN(outfile, fileRdr, unzipSize)
	if err != nil {
		return err
	}
	return nil
}
