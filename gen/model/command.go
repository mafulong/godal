package model

import (
	"fmt"
	"github.com/mafulong/godal/utils"
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
	fmt.Println(utils.ToJSON(args))
	fmt.Println(c.StringSlice("abc"))
	return nil
}

var Command = &cli.Command{
	Name:      "model",
	Usage:     "generate model directory",
	ArgsUsage: "[tableSQLs...]",
	Flags:     flags,
	Action:    action,
}
