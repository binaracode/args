package main

import (
	"fmt"

	"github.com/buglk/args"
)

type Command2 struct {
	flagArg   bool
	stringArg string
}

func (c *Command2) ListArgs() []args.Arg {
	var a []args.Arg

	/* a mandatory flag argument */
	a = append(a, args.Flag(&c.flagArg, "flag", "example flag arg"))

	/* a string argument */
	a = append(a, args.String(&c.stringArg, "str", "example string arg"))

	return a
}

func (c *Command2) IsDefault() bool {
	return false
}

func (c *Command2) Name() string {
	return "two"
}

func (c *Command2) Help() string {
	return "command-2 help"
}

func (c *Command2) Run() (int, error) {
	fmt.Printf("cmd2->flagArg: %v\n", c.flagArg)
	fmt.Printf("cmd2->stringArg: %s\n", c.stringArg)
	return 0, nil
}
