package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Age    int    `cli:"age" usage:"your age"`
	Gender string `cli:"g,gender" usage:"your gender" dft:"male"`
}

func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.Age < 0 || argv.Age > 200 {
		return fmt.Errorf("age %d out of range\n", argv.Age)
	}
	if argv.Gender != "male" && argv.Gender != "female" {
		return fmt.Errorf("invaild gender %s\n", ctx.Color().Yellow(argv.Gender))
	}
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		return nil
	})
}
