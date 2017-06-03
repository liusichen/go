package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	ID uint8 `cli:"*id" usage:"this is a required parameter, note the *"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("%d\n", argv.ID)
		return nil
	})
}
