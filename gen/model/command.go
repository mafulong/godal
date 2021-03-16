package model

import (
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:  "database",
		Usage: "database of the create table SQL",
	},
	&cli.StringSliceFlag{
		Name:  "abc",
		Usage: "ab",
	},
}

func action(c *cli.Context) error {
	args := c.Args()
	for i := 0; i < args.Len(); i += 1 {
		err := gen(args.Get(i))
		if err != nil {
			return err
		}
	}
	return nil
}

var Command = &cli.Command{
	Name:      "model",
	Usage:     "generate model directory",
	ArgsUsage: "[tableSQLs...]",
	Flags:     flags,
	Action:    action,
}
