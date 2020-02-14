package main

import (
	"flag"
	"log"

	"github.com/locona/github-release-qadoc/pkg/github"
	"github.com/locona/github-release-qadoc/pkg/markdown"
)

var (
	issueNumber  int
	organization string
	repo         string
	ci           bool
)

func main() {
	flag.IntVar(&issueNumber, "issue", 0, "GitHub Issue Number.")
	flag.StringVar(&organization, "organization", "", "GitHub Organization.")
	flag.StringVar(&repo, "repo", "", "GitHub Repository.")
	flag.BoolVar(&ci, "ci", false, "Use CI")
	flag.Parse()

	var opt *github.Option
	if issueNumber == 0 {
		log.Fatal("Required Issue Number(--issue)")
	}
	if ci {
		if organization == "" {
			log.Fatal("Required GitHub Organization(--organization)")
		}

		if repo == "" {
			log.Fatal("Required GitHub Repository(--repo)")
		}

		opt = &github.Option{
			Organization: organization,
			Repository:   repo,
		}
	}

	cli, err := github.New(opt)
	if err != nil {
		panic(err)
	}

	release, err := cli.LatestRelease()
	if err != nil {
		panic(err)
	}

	res := markdown.List2TodoList(*release.Body)
	_, err = cli.IssueComment(issueNumber, res)
	if err != nil {
		panic(err)
	}
}
