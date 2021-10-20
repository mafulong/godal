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

func run(args []string) error {
	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	_ = run(os.Args)
}
