package model

import (
	"github.com/urfave/cli/v2"
)

const (
	keyDatabase = "database"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:  keyDatabase,
		Usage: "optional. database of the create table SQL",
	},
}

// Args cms args
type Args struct {
	Database string
}

var cmdArgs = Args{}

func action(c *cli.Context) error {
	args := c.Args()
	if len(c.String(keyDatabase)) > 0 {
		cmdArgs.Database = c.String(keyDatabase)
	}
	for i := 0; i < args.Len(); i ++{
		err := gen(args.Get(i))
		if err != nil {
			return err
		}
	}
	return nil
}

// Command for cmd
var Command = &cli.Command{
	Name:      "model",
	Usage:     "generate model directory",
	ArgsUsage: "[tableSQLs...]",
	Flags:     flags,
	Action:    action,
}
