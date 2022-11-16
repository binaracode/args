package args

type FlagArg struct {
	Arg
	value *bool
}

func Flag(value *bool, name string, help string) *FlagArg {
	sa := &FlagArg{
		Arg:   newArgBase(name, help),
		value: value,
	}

	return sa
}

func (sa *FlagArg) Parse(args []string, index int) (int, error) {
	ba := sa.Arg.(*ArgBase)
	i, err := ba.Parse(args, index)
	if i == -1 || err != nil {
		return i, err
	}

	*sa.value = true
	ba.set = true
	return index, nil
}
