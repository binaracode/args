package args

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SubCmd interface {
	Run() (int, error)
	Name() string
	Help() string
	IsDefault() bool
	ListArgs() []Arg
}

type CommandCollection struct {
	commands []SubCmd

	Usage           func(string) string
	ArgError        func(error) error
	CommandArgError func(SubCmd, string) error
	UnknownArgs     func(SubCmd, []string) error
	CommandUsage    func(SubCmd) string
}

func NewCollection() *CommandCollection {
	cc := &CommandCollection{}
	cc.ArgError = cc.argError
	cc.UnknownArgs = cc.unknownArgs
	cc.Usage = cc.usage
	cc.CommandUsage = cc.commandUsage
	cc.CommandArgError = cc.cmdArgError
	return cc
}

func (c *CommandCollection) usage(msg string) string {
	_, exe := filepath.Split(os.Args[0])
	if msg != "" {
		msg += "\n"
	}
	str := fmt.Sprintf("%sUsage:%s [command] [options]\nCommands:\n", msg, exe)
	for _, cmd := range c.commands {
		str += fmt.Sprintf(" %-10s - %s\n", cmd.Name(), cmd.Help())
	}

	h := ""
	for i, a := range help {
		if i > 0 {
			h += ", "
		}
		h += a
	}
	str += fmt.Sprintf("\nUse \"%s [command] %s\" for more information about a given command\n", exe, h)
	return str
}

func (c *CommandCollection) commandUsage(cmd SubCmd) string {
	_, exe := filepath.Split(os.Args[0])
	str := fmt.Sprintf("Summary:\n %s\nUsage:\n %s %s [options]\nOptions:\n", cmd.Help(), exe, cmd.Name())
	argList := cmd.ListArgs()
	argc := len(argList)
	for i := 0; i < argc; i++ {
		if i > 0 {
			str += "\n"
		}
		strs := strings.Split(argList[i].ToString(), "\n")
		for _, s := range strs {
			str += fmt.Sprintf("   %s\n", s)
		}
	}
	return str
}

/* defult error handling */
func (c *CommandCollection) argError(err error) error {
	return fmt.Errorf("%s", c.Usage(err.Error()))
}

func (c *CommandCollection) unknownArgs(sc SubCmd, args []string) error {
	ar := ""
	for i, s := range args {
		if i > 0 {
			ar += ", "
		}
		ar += s
	}
	return fmt.Errorf("unknown argument(s): %s", ar)
}

func (c *CommandCollection) cmdArgError(cmd SubCmd, msg string) error {
	err := "Argument error:\n " + msg
	err += "\n\n" + c.CommandUsage(cmd)
	return fmt.Errorf(err)
}

func (c *CommandCollection) Add(cmd SubCmd) {
	for _, c := range c.commands {
		if c == cmd {
			return
		}
		if c.Name() == cmd.Name() {
			panic(fmt.Sprintf("duplicate command: %s", c.Name()))
		}
	}
	c.commands = append(c.commands, cmd)
}

func (c *CommandCollection) isHelp(arg *string) bool {
	for _, a := range help {
		if a == *arg {
			return true
		}
	}
	return false
}

func (c *CommandCollection) runDefault(args *[]string, index int) (int, error) {
	/* find a default command */
	for _, cmd := range c.commands {
		if cmd.IsDefault() {
			return c.invoke(cmd, args, index)
		}
	}

	return 1, c.ArgError(fmt.Errorf("missing sub command"))
}

func (c *CommandCollection) Run() (int, error) {
	args, err := c.prepareArgs()
	if err != nil {
		return 1, c.ArgError(err)
	}

	argc := len(args)

	/* call for help */
	if argc == 1 {
		if c.isHelp(&args[0]) {
			return 1, fmt.Errorf(c.Usage(""))
		}
	}

	if argc == 0 || (argc > 1 && args[0][0] == argPrefix) {
		/* no args (len ==0) or first arg is a flag (-/--) or option - try default command */
		return c.runDefault(&args, 0)
	}

	subCmd := args[0]
	for _, cmd := range c.commands {
		if subCmd == cmd.Name() {
			return c.invoke(cmd, &args, 1)
		}
	}

	return 1, c.ArgError(fmt.Errorf("sub command '%s' was not found", subCmd))
}

func (c *CommandCollection) parseAndValidate(cmd SubCmd, args *[]string, index int) error {

	argc := len(*args)
	if argc == index+1 {
		if c.isHelp(&(*args)[index]) {
			return fmt.Errorf(c.CommandUsage(cmd))
		}
	}

	var errs []error
	var unused []string
	argList := cmd.ListArgs()
	for i := index; i < argc; i++ {
		unknown := true
		for j := 0; j < len(argList); j++ {
			next, err := argList[j].Parse(*args, i)
			if err != nil {
				errs = append(errs, err)
				unknown = false
				break
			} else {
				if next != -1 {
					i = next
					unknown = false
					break
				}
			}
		}
		if unknown {
			unused = append(unused, (*args)[i])
		}
	}

	if len(errs) > 0 {
		return c.CommandArgError(cmd, c.toString(errs))
	}

	if len(unused) > 0 {
		err := c.UnknownArgs(cmd, unused)
		if err != nil {
			return c.CommandArgError(cmd, err.Error())
		}
	}
	var invalids []error
	for i := 0; i < len(argList); i++ {
		err := argList[i].Validate()
		if err != nil {
			invalids = append(invalids, err)
		}
	}
	if len(invalids) > 0 {
		return c.CommandArgError(cmd, c.toString(invalids))
	}
	return nil
}

func (c *CommandCollection) toString(errs []error) string {
	if len(errs) == 1 {
		return errs[0].Error()
	}
	var err string
	for i, e := range errs {
		if i > 0 {
			err += "\n "
		}
		err += e.Error()
	}
	return err
}

func (c *CommandCollection) invoke(cmd SubCmd, args *[]string, index int) (int, error) {
	err := c.parseAndValidate(cmd, args, index)
	if err != nil {
		return 1, err
	}
	return cmd.Run()
}
