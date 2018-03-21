package main

import (
	"context"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/tcnksm/go-gitconfig"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var (
	username = "Ken2mer"
)

var githubCommand = cli.Command{
	Name:   "github",
	Usage:  "github",
	Action: githubCmd,
}

func githubCmd(c *cli.Context) error {
	ctx := context.Background()
	client := oauth2Client(ctx)

	user, err := getUser(ctx, client)
	if err != nil {
		return err
	}
	follows, err := getFollowingUsers(ctx, client)
	if err != nil {
		return err
	}
	code, err := getCodeSearchResult(ctx, client)
	if err != nil {
		return err
	}
	events, err := getEvents(ctx, client)
	if err != nil {
		return err
	}

	dumpUser(user)
	dumpFollowUsers(follows)
	dumpCodeSearchResult(code)
	dumpEvents(events)

	return nil
}

// cf. https://github.com/mackerelio/mkr/blob/af4f89ae6fac2290b9fe642de37f84a25de67d62/plugin/github.go

// Get github client having github token.
func oauth2Client(ctx context.Context) *github.Client {
	var oauthClient *http.Client
	if token := getGithubToken(); token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: getGithubToken()},
		)
		oauthClient = oauth2.NewClient(ctx, ts)
	}
	return github.NewClient(oauthClient)
}

// Get github token from environment variables, or github.token in gitconfig file
func getGithubToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		return token
	}
	token, _ = gitconfig.GithubToken()
	return token
}
