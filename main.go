package main

import (
	"flag"
	"strconv"

	"github.com/locona/github-release-qadoc/pkg/github"
	"github.com/locona/github-release-qadoc/pkg/markdown"
)

func main() {
	cli := github.New()
	flag.Parse()
	args := flag.Args()
	issueNumber, err := strconv.Atoi(args[0])
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
