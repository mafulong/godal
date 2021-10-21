package main

import (
	"github.com/mafulong/godal/gen"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
)

var commands = []*cli.Command{
	gen.Command,
}

var app = &cli.App{
	Name:  "godal",
	Usage: "generate file",
	Action: func(context *cli.Context) error {
		return nil
	},
	Commands: commands,
}

func main() {
	_ = app.Run(os.Args)
}
