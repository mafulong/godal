package main

import (
	"github.com/mafulong/godal/gen"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

var commands = []*cli.Command{
	gen.Command,
}

func action(c *cli.Context) error {
	return nil
}

var app = &cli.App{
	Name:     "godal",
	Usage:    "generate file",
	Action:   action,
	Commands: commands,
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
