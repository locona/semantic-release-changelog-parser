package github

import (
	"context"

	"github.com/google/go-github/v29/github"
)

func (cli *Client) LatestRelease() (*github.RepositoryRelease, error) {
	ctx := context.Background()
	release, _, err := cli.svc.Repositories.GetLatestRelease(
		ctx,
		cli.Config.Organization,
		cli.Config.Repository,
	)

	if err != nil {
		return nil, err
	}

	return release, nil
}

func (cli *Client) IssueComment(issueNumber int, body string) (*github.IssueComment, error) {
	ctx := context.Background()
	release, _, err := cli.svc.Issues.CreateComment(
		ctx,
		cli.Config.Organization,
		cli.Config.Repository,
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
