package main

import (
	"fmt"

	"github.com/buglk/args"
)

type Command1 struct {
	intArg    int
	stringArg string
}

func (c *Command1) ListArgs() []args.Arg {
	var a []args.Arg

	/* a mandatory integer argument */
	a = append(a, args.Integer(&c.intArg, "int-arg", "example integer arg").ShortCode('i').Mandatory())

	/* a string argument */
	a = append(a, args.String(&c.stringArg, "str", "example string arg"))

	return a
}

func (c *Command1) IsDefault() bool {
	return false
}

func (c *Command1) Name() string {
	return "one"
}

func (c *Command1) Run() (int, error) {
	fmt.Printf("cmd1->intArg: %d\n", c.intArg)
	fmt.Printf("cmd1->stringArg: %s\n", c.stringArg)
	return 0, nil
}

func (c *Command1) Help() string {
	return "command-1 help"
}
