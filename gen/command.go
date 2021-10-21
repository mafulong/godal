package gen

import (
	"github.com/mafulong/godal/gen/model"
	"github.com/urfave/cli/v2"
)

var subCommands = []*cli.Command{
	model.Command,
}

// Command gen command
var Command = &cli.Command{
	Name:        "gen",
	Usage:       "generators",
	Subcommands: subCommands,
}
