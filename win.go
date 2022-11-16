//go:build (windows && !linux_style) || (windows_style && linux)

package args

import (
	"os"
)

var help = []string{"/?", "/h", "/help"}

const argPrefix byte = '/'
const argPrefixLong string = "/"

func (a *ArgBase) matched(arg string) bool {
	if arg[0] == '/' {
		arg = arg[1:]
	}
	if a.name == arg {
		return true
	}
	if len(arg) == 1 && a.short == arg[0] {
		return true
	}
	return false
}

func (c *CommandCollection) prepareArgs() ([]string, error) {
	return os.Args[1:], nil
}
