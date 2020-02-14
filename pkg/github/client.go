package github

import (
	"context"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v29/github"

	"github.com/locona/github-release-qadoc/pkg/gitconfig"
)

type Client struct {
	svc    *github.Client
	Config *gitconfig.RemoteConfig
}

type Option struct {
	Organization string
	Repository   string
}

func New(opt *Option) (*Client, error) {
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)

	if opt != nil {
	}

	var remoteConfig *gitconfig.RemoteConfig
	if opt == nil {
		gc, err := gitconfig.Config()
		if err != nil {
			return nil, err
		}

		remoteConfig = gc.RemoteConfig
	} else {
		remoteConfig = &gitconfig.RemoteConfig{
			Organization: opt.Organization,
			Repository:   opt.Repository,
		}

	}

	return &Client{
		svc:    githubClient,
		Config: remoteConfig,
	}, nil
}
