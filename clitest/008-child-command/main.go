package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

type rootT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}
type childT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name"`
}

var help = cli.HelpCommand("display help information")
var root = &cli.Command{
	Desc: "this is root command",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}

var child = &cli.Command{
	Name: "child",
	Desc: "this is child command",
	Argv: func() interface{} { return new(childT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childT)
		ctx.String("Hello, child command, I am %s\n", argv.Name)
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(child),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
