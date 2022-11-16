package args

import "fmt"

type StringArg struct {
	Arg
	value *string
}

func String(value *string, name string, help string) *StringArg {
	sa := &StringArg{
		Arg:   newArgBase(name, help),
		value: value,
	}

	return sa
}

func (sa *StringArg) ShortCode(code byte) *StringArg {
	ba := sa.Arg.(*ArgBase)
	ba.ShortCode(code)
	return sa
}

func (sa *StringArg) Mandatory() *StringArg {
	ba := sa.Arg.(*ArgBase)
	ba.Mandatory()
	return sa
}

func (sa *StringArg) Parse(args []string, index int) (int, error) {
	ba := sa.Arg.(*ArgBase)
	i, err := ba.Parse(args, index)
	if i == -1 || err != nil {
		return i, err
	}

	if len(args) < index+2 {
		return -1, fmt.Errorf("string argument '%s' require a value", ba.name)
	}

	*sa.value = args[index+1]
	ba.set = true
	return index + 1, nil
}

func (sa *StringArg) ToString() string {
	ba := sa.Arg.(*ArgBase)
	return ba.ToString()
}
