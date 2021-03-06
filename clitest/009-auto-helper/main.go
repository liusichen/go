package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	Help bool `cli:"h,help" usage:"show help"`
}

func (argv *argT) AutoHelp() bool {
	return argv.Help
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		return nil
	})
}
