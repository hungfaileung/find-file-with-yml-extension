package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func showFiles(ctx context.Context, user string, repo string, path string, client *github.Client, ext string) {
	_, directoryContent, _, err := client.Repositories.GetContents(ctx, user, repo, path, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, dc := range directoryContent {
		if *dc.Type == "file" {
			if filepath.Ext(*dc.Path) == ext {
				fmt.Printf("%-10s | %-20s | %-80s\n", user, repo, *dc.Path)
			}
		} else if *dc.Type == "dir" {
			showFiles(ctx, user, repo, *dc.Path, client, ext)
		} else {
			fmt.Println(*dc.Name, *dc.Path, *dc.Type)
		}
	}
}

func main() {

	ext := ".html"

	token := os.Getenv("GITHUB_PAT")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	showFiles(ctx, "hungfaileung", "01-build-a-simple-web-server-with-golang", "", client, ext)
	showFiles(ctx, "hungfaileung", "01-build-a-simple-web-server-with-golang", "", client, ".yml")

}
