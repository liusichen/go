package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	Friends []string `cli:"F" usage:"my friends,you can use comma(,) to conbine the friends "`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		return nil
	})
}
