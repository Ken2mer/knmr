package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Ken2mer/knmr/logger"
	"github.com/urfave/cli"
)

var (
	placeID         int64  = 1118550    // Yokohama
	slug            string = "engineer" // for GetListTweetsBySlug()
	ownerScreenName string = "Ken2mer"  // for GetListTweetsBySlug()
	queryString     string = "lang:ja"  // for GetSearch()
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
	dumpListTweets(api)
	return nil
}

func dumpTrendResp(api *anaconda.TwitterApi) {
	trendResp, err := api.GetTrendsByPlace(placeID, nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	logger.DumpJSON(trendResp)
}

func dumpListTweets(api *anaconda.TwitterApi) {
	tweets, err := api.GetListTweetsBySlug(slug, ownerScreenName, true, nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	logger.DumpJSON(tweets)
}

func dmupTimeline(api *anaconda.TwitterApi) {
	timeline, err := api.GetHomeTimeline(nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	logger.DumpJSON(timeline)
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
	searchResponse, err := api.GetSearch(queryString, nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	for {
		logger.DumpJSON(searchResponse)

		searchResponse, err = searchResponse.GetNext(api)
		if searchResponse.Statuses == nil {
			break
		}
	}
}
