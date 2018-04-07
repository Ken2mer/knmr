package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/urfave/cli"
)

type lbClient struct {
	bot *linebot.Client
}

var linebotCommand = cli.Command{
	Name:   "linebot",
	Usage:  "linebot",
	Action: linebotCmd,
}

func linebotCmd(ctx *cli.Context) error {
	// Authentication info
	// var channelSecret string = "XXXXXX"
	// var channelToken  string = "XXXXXX"

	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return err
	}

	c := lbClient{bot: bot}
	return c.pushMessage()
}

// cf. https://github.com/line/line-bot-sdk-go#create-message
func (c *lbClient) pushMessage() error {
	var bot *linebot.Client = c.bot
	var messages []linebot.Message

	leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")
	rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")
	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
	message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

	messages = append(messages, message)

	// var userID string = "XXXXXX"
	_, err := bot.PushMessage(userID, messages...).Do()
	if err != nil {
		return err
	}
	return nil
}
