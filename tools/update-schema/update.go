package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/requests/repos"
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
	if tag == "latest" {
		tag, err = getNewTag(ctx, client)
		if err != nil {
			log.Fatal(err)
		}
		if tag == "" {
			log.Println("nothing to update to")
			return
		}
	}

	currentVersion, err := getCurrentVersion(workingDir)
	if err != nil {
		log.Fatal(err)
	}
	if currentVersion == tag && !force {
		fmt.Println("already up to date")
		return
	}

	err = downloadFromRelease(tag, filepath.Join(workingDir, "api.github.com.json"))
	if err != nil {
		log.Fatal(err)
	}
	err = setCurrentVersion(workingDir, tag)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFromRelease(tag, destination string) error {
	pattern := `https://raw.githubusercontent.com/github/rest-api-description/%s/descriptions/api.github.com/api.github.com.json`
	resp, err := http.Get(fmt.Sprintf(pattern, tag))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("it's not OK")
	}
	outfile, err := os.Create(destination)
	if err != nil {
		return err
	}
	_, err = io.Copy(outfile, resp.Body)
	return err
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

func getNewTag(ctx context.Context, client octo.Client) (string, error) {
	releaseTag, err := latestReleaseTag(ctx, client)
	if err != nil {
		return "", err
	}
	tagTag, err := newestSemverTag(ctx, client)
	if err != nil {
		return "", err
	}
	tag, err := higherValidSemver(releaseTag, tagTag)
	if err != nil && err != errTwoInvalidSemvers {
		return "", err
	}
	return tag, nil
}

func latestReleaseTag(ctx context.Context, client octo.Client) (string, error) {
	resp, err := client.Repos().GetLatestRelease(ctx, &repos.GetLatestReleaseReq{
		Owner: "github",
		Repo:  "rest-api-description",
	})
	if err != nil {
		return "", err
	}
	return resp.Data.TagName, nil
}

func newestSemverTag(ctx context.Context, client octo.Client) (string, error) {
	resp, err := client.Repos().ListTags(ctx, &repos.ListTagsReq{
		Owner: "github",
		Repo:  "rest-api-description",
	})
	if err != nil {
		return "", err
	}
	result := "v0-phony"
	for _, tag := range resp.Data {
		result, err = higherValidSemver(result, tag.Name)
		if err != nil {
			return "", err
		}
	}
	if result == "v0-phony" {
		result = ""
	}
	return result, nil
}

var errTwoInvalidSemvers = fmt.Errorf("can't compare two invalid semvers")

func higherValidSemver(a, b string) (string, error) {
	verA, err := semver.NewVersion(a)
	if err != nil && err != semver.ErrInvalidSemVer {
		return "", err
	}
	verB, err := semver.NewVersion(b)
	if err != nil && err != semver.ErrInvalidSemVer {
		return "", err
	}
	if verA == nil && verB == nil {
		return "", errTwoInvalidSemvers
	}
	if verA == nil {
		return b, nil
	}
	if verB == nil {
		return a, nil
	}
	if verB.GreaterThan(verA) {
		return b, nil
	}
	return a, nil
}
