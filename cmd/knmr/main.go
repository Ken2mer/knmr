package main

import (
	"fmt"
	"os"

	"github.com/Ken2mer/knmr/logger"
	"github.com/urfave/cli"
)

const (
	name    = "knmr"
	version = "0.0.0"
)

func init() {
	logger.Infof("%s version: %s\n", name, version)

	for _, f := range os.Args {
		if f == "-v" || f == "--verbose" || f == "-verbose" {
			logger.Verbose = true
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Usage = "knmr utility"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug output in logs",
		},
	}

	app.Commands = []cli.Command{
		agentCommand,
	}

	var debugEnabled bool

	app.Before = func(context *cli.Context) error {
		debugEnabled = context.GlobalBool("debug")
		if debugEnabled {
			logger.Verbose = true
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
