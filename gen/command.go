package gen

import (
	"github.com/mafulong/godal/gen/model"
	"github.com/urfave/cli/v2"
)

var subCommands = []*cli.Command{
	model.Command,
}

func action(_ *cli.Context) error {
	return nil
}

//gen command
var Command = &cli.Command{
	Name:        "gen",
	Usage:       "generators",
	Action:      action,
	Subcommands: subCommands,
}
