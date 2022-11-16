package main

import (
	"fmt"
	"os"

	"github.com/buglk/args"
)

func main() {
	cc := args.NewCollection()
	cc.Add(&Command1{})
	cc.Add(&Command2{})
	cc.Add(NewCmd3())

	exit, err := cc.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(exit)
}
