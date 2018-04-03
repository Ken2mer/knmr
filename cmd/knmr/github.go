package main

import (
	"context"
	"net/http"
	"os"

	"github.com/Ken2mer/knmr/logger"
	"github.com/google/go-github/github"
	"github.com/tcnksm/go-gitconfig"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

var (
	username string = "Ken2mer"
	reponame string = "knmr"
)

var githubCommand = cli.Command{
	Name:   "github",
	Usage:  "github",
	Action: githubCmd,
}

func githubCmd(c *cli.Context) error {
	ctx := context.Background()
	client := oauth2Client(ctx)

	serve()
	// return subscription(ctx, client)
	// return events(ctx, client)
	return code(ctx, client)
	// return follows(ctx, client)
	// return user(ctx, client)
}

func serve() {
	s := gitHubEventMonitor{
		webhookSecretKey: []byte(getGithubToken()),
	}
	http.HandleFunc("/", s.serveHTTP)
	logger.Error(http.ListenAndServe(":12345", nil))
}

func subscription(ctx context.Context, client *github.Client) error {
	subscription, err := getRepositorySubscription(ctx, client)
	if err != nil {
		return err
	}
	logger.DumpJSON(subscription)
	return nil
}

func events(ctx context.Context, client *github.Client) error {
	events, err := getEvents(ctx, client)
	if err != nil {
		return err
	}
	dumpEvents(events)
	return nil
}

func code(ctx context.Context, client *github.Client) error {
	code, err := getCodeSearchResult(ctx, client)
	if err != nil {
		return err
	}
	logger.DumpJSON(code)
	return nil
}

func follows(ctx context.Context, client *github.Client) error {
	follows, err := getFollowingUsers(ctx, client)
	if err != nil {
		return err
	}
	dumpFollowUsers(follows)
	return nil
}

func user(ctx context.Context, client *github.Client) error {
	user, err := getUser(ctx, client)
	if err != nil {
		return err
	}
	dumpUser(user)
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
