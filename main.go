package main

import (
	"github.com/mkideal/cli"
	"os"
)

type CmdArgs struct {
	// []bool, []int, []float32, ... supported too.
	PrimaryKeys []string `cli:"pk" usage:"[required, array]primaryKey"`
	TableName string `cli:"table_name" usage:"[required, string]table_name"`
}

var CmdArgsIns CmdArgs

func main() {
	os.Exit(cli.Run(new(CmdArgs), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		args := ctx.Argv().(*CmdArgs)
		//utils.ToJSON(args)
		return nil
	}))
}
