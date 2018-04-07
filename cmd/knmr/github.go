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

type ghClient struct {
	ctx context.Context
	client *github.Client
}

var githubCommand = cli.Command{
	Name:   "github",
	Usage:  "github",
	Action: githubCmd,
}

func githubCmd(ctx *cli.Context) error {
	context := context.Background()
	client := oauth2Client(context)

	c := ghClient{ctx: context, client: client}

	serve()
	// return c.subscription()
	// return c.events()
	return c.code()
	// return c.follows()
	// return c.user()
}

func serve() {
	s := gitHubEventMonitor{
		webhookSecretKey: []byte(getGithubToken()),
	}
	http.HandleFunc("/payload", s.serveHTTP)
	logger.Error(http.ListenAndServe(":12345", nil))
}

func (c *ghClient) subscription() error {
	subscription, err := getRepositorySubscription(c.ctx, c.client)
	if err != nil {
		return err
	}
	logger.DumpJSON(subscription)
	return nil
}

func (c *ghClient) events() error {
	events, err := getEvents(c.ctx, c.client)
	if err != nil {
		return err
	}
	dumpEvents(events)
	return nil
}

func (c *ghClient) code() error {
	code, err := getCodeSearchResult(c.ctx, c.client)
	if err != nil {
		return err
	}
	logger.DumpJSON(code)
	return nil
}

func (c *ghClient) follows() error {
	follows, err := getFollowingUsers(c.ctx, c.client)
	if err != nil {
		return err
	}
	dumpFollowUsers(follows)
	return nil
}

func (c *ghClient) user() error {
	user, err := getUser(c.ctx, c.client)
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
