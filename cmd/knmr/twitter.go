package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

var twitterCommand = cli.Command{
	Name:   "twitter",
	Usage:  "twitter",
	Action: twitterCmd,
}

func twitterCmd(c *cli.Context) error {
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)

	tweets, err := api.GetListTweetsBySlug("engineer", "Ken2mer", true, nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, tweet := range tweets {
		fmt.Printf("%d: \n", i)
		fmt.Printf("%s\n", tweet.CreatedAt)
		fmt.Printf("%s wrote: \n", tweet.User.Name)
		fmt.Printf("%s\n\n", tweet.FullText)
	}

	return nil
}

func dmupTimeline(api *anaconda.TwitterApi) {
	timeline, err := api.GetHomeTimeline(nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, tweet := range timeline {
		fmt.Printf("%d: \n", i)
		fmt.Printf("%s\n", tweet.CreatedAt)
		fmt.Printf("%s wrote: \n", tweet.User.Name)
		fmt.Printf("%s\n\n", tweet.FullText)
	}
}

func dumpFriends(api *anaconda.TwitterApi) {
	cursor, err := api.GetFriendsList(nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, user := range cursor.Users {
		fmt.Printf("%d: \n", i)
		fmt.Printf("user: %s\n", user.Name)
	}
}

func dumpFollowers(api *anaconda.TwitterApi) {
	cursor, err := api.GetFollowersList(nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, user := range cursor.Users {
		fmt.Printf("%d: \n", i)
		fmt.Printf("user: %s\n", user.Name)
	}
}

func dumpFavorites(api *anaconda.TwitterApi) {
	favorites, err := api.GetFavorites(nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, tweet := range favorites {
		fmt.Printf("%d: \n", i)
		fmt.Printf("%s wrote: \n", tweet.User.Name)
		fmt.Printf("%s\n\n", tweet.FullText)
	}
}

func dumpSearchResult(api *anaconda.TwitterApi) {
	searchResult, err := api.GetSearch("golang", nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for i, tweet := range searchResult.Statuses {
		fmt.Printf("%d: \n", i)
		fmt.Printf("%s wrote: \n", tweet.User.Name)
		fmt.Printf("%s\n\n", tweet.FullText)
	}
}
