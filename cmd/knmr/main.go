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

func main() {
	logger.Infof("%s version: %s\n", name, version)

	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Usage = "knmr"

	app.Commands = commands
	app.Flags = flags
	app.Before = before

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
