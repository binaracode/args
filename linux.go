//go:build (linux && !windows_style) || (linux_style && windows)

package args

import (
	"fmt"
	"os"
	"strings"
)

var help = []string{"-?", "--help"}

const argPrefix byte = '-'
const argPrefixLong string = "--"

func (a *ArgBase) matched(arg string) bool {
	delims := 0
	for arg[delims] == '-' {
		delims++
	}
	arg = arg[delims:]
	if delims == 2 && a.name == arg {
		return true
	}
	if delims == 1 && len(arg) == 1 && a.short == arg[0] {
		return true
	}
	return false
}

func (c *CommandCollection) prepareArgs() ([]string, error) {
	var temp []string
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "--") {
			temp = append(temp, arg)
		} else if arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				temp = append(temp, fmt.Sprintf("-%c", arg[i]))
			}
		} else if arg[0] == '/' {
			return temp, fmt.Errorf("argument '%s' not accepted", arg)
		} else {
			temp = append(temp, arg)
		}
	}

	return temp, nil
}
