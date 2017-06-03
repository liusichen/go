package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Username string `cli:"u,username" usage:"github account" prompt:"type github account"`
	Password string `cli:"p,password" usage:"password of the github account" prompt:"type the password"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("username = %s, password = %s\n", argv.Username, argv.Password)
		return nil
	})
}
