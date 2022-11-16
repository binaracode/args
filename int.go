package args

import (
	"fmt"
	"strconv"
)

type IntegerArg struct {
	Arg
	value *int
}

func Integer(value *int, name string, help string) *IntegerArg {
	sa := &IntegerArg{
		Arg:   newArgBase(name, help),
		value: value,
	}

	return sa
}

func (ia *IntegerArg) ShortCode(code byte) *IntegerArg {
	ba := ia.Arg.(*ArgBase)
	ba.ShortCode(code)
	return ia
}

func (ia *IntegerArg) Mandatory() *IntegerArg {
	ba := ia.Arg.(*ArgBase)
	ba.Mandatory()
	return ia
}

func (ia *IntegerArg) Parse(args []string, index int) (int, error) {
	ba := ia.Arg.(*ArgBase)
	i, err := ba.Parse(args, index)
	if i == -1 || err != nil {
		return i, err
	}

	if len(args) < index+2 {
		return -1, fmt.Errorf("integer argument '%s' require a value", ba.name)
	}

	i64, err := strconv.ParseInt(args[index+1], 10, 0)
	if err != nil {
		return -1, fmt.Errorf("%s value format error: %s", ba.name, err.Error())
	}

	*ia.value = int(i64)
	ba.set = true
	return index + 1, nil
}
