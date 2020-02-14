package github

import (
	"context"

	"github.com/google/go-github/v29/github"

	"github.com/locona/github-release-qadoc/pkg/gitconfig"
)

func (cli *Client) LatestRelease() (*github.RepositoryRelease, error) {
	ctx := context.Background()
	gc, err := gitconfig.Config()
	if err != nil {
		return nil, err
	}

	release, _, err := cli.svc.Repositories.GetLatestRelease(
		ctx,
		gc.RemoteConfig.Organization,
		gc.RemoteConfig.Repository,
	)

	if err != nil {
		return nil, err
	}

	return release, nil
}

func (cli *Client) IssueComment(issueNumber int, body string) (*github.IssueComment, error) {
	ctx := context.Background()
	gc, err := gitconfig.Config()
	if err != nil {
		return nil, err
	}

	release, _, err := cli.svc.Issues.CreateComment(
		ctx,
		gc.RemoteConfig.Organization,
		gc.RemoteConfig.Repository,
		issueNumber,
		&github.IssueComment{
			Body: &body,
		},
	)

	if err != nil {
		return nil, err
	}

	return release, nil
}
