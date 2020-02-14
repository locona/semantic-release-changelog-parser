package github

import (
	"context"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v29/github"
)

type Client struct {
	svc *github.Client
}

func New() *Client {
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	return &Client{svc: githubClient}
}
