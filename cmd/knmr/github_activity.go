package main

import (
	"context"
	"fmt"

	"github.com/Ken2mer/knmr/logger"
	"github.com/google/go-github/github"
)

func getEvents(ctx context.Context, client *github.Client) ([]*github.Event, error) {
	events, _, err := client.Activity.ListEventsReceivedByUser(ctx, username, true, nil)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		logger.Errorf("\nerror: %v\n", err)
		return nil, err
	}
	return events, nil
}

func dumpEvents(events []*github.Event) {
	for i, event := range events {
		logger.Debugf("%d\n", i)
		fmt.Printf("%s\n\n", event.GetRawPayload())
	}
}
