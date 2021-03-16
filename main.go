package main

import (
	"fmt"
	"github.com/mafulong/godal/utils"
	"github.com/mkideal/cli"
)

type CmdArgs struct {
	cli.Helper
	TableName   string   `cli:"*table_name" usage:"table_name"`
	PrimaryKeys []string `cli:"*pk" usage:"primaryKey"`
}

var CmdArgsIns CmdArgs

func main() {
	cli.Run(new(CmdArgs), func(ctx *cli.Context) error {
		//ctx.JSONln(ctx.Argv())
		args := ctx.Argv().(*CmdArgs)
		fmt.Println(utils.ToJSON(args))
		return nil
	})
}
