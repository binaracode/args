package args

import (
	"fmt"
)

type Arg interface {
	Parse(args []string, index int) (int, error)
	Validate() error
	ToString() string
}

type ArgBase struct {
	name     string
	short    byte
	required bool
	help     string
	set      bool
}

func newArgBase(name string, help string) Arg {
	a := &ArgBase{
		name:     name,
		required: false,
		help:     help,
		set:      false,
	}
	return a
}

func (a *ArgBase) Validate() error {
	if a.required && !a.set {
		return fmt.Errorf("'%s' is required", a.name)
	}
	return nil
}

func (a *ArgBase) Parse(args []string, index int) (int, error) {
	if a.matched(args[index]) {
		if a.set {
			return -1, fmt.Errorf("'%s' already set", a.name)
		}
		return index, nil
	}
	return -1, nil
}

func (a *ArgBase) ShortCode(code byte) *ArgBase {
	a.short = code
	return a
}

func (a *ArgBase) Mandatory() *ArgBase {
	a.required = true
	return a
}

func (a *ArgBase) ToString() string {
	str := ""
	if a.short != 0 {
		str += fmt.Sprintf("%c%c, ", argPrefix, a.short)
	}
	reqflag := ""
	if a.required {
		reqflag = "(required)"
	}
	return fmt.Sprintf("%s%s%s: %s\n   %s", str, argPrefixLong, a.name, reqflag, a.help)
}
