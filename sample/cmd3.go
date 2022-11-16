package main

import (
	"fmt"

	"github.com/buglk/args"
)

type Command3 struct {
	intArg    int
	stringArg string
}

func NewCmd3() *Command3 {
	c3 := &Command3{
		stringArg: "default-value",
	}
	return c3
}

func (c *Command3) ListArgs() []args.Arg {
	var a []args.Arg

	/* a mandatory flag argument */
	a = append(a, args.Integer(&c.intArg, "int", "cmd3 -> example flag arg").ShortCode('i'))

	/* a string argument */
	a = append(a, args.String(&c.stringArg, "str", "cmd3 -> example string arg - default value is: 'default-value'"))

	return a
}

func (c *Command3) IsDefault() bool {
	return false
}

func (c *Command3) Name() string {
	return "three"
}

func (c *Command3) Help() string {
	return "command-3 help"
}

func (c *Command3) Run() (int, error) {
	fmt.Printf("cmd3->intArg: %v\n", c.intArg)
	fmt.Printf("cmd3->stringArg: %s\n", c.stringArg)
	return 0, nil
}
