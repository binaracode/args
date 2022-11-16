# args

This is a simple go library that supports sub commands (as in flag sets). The sub command is hooked to a function.

# sample

main code:
```
	// create a command collection
	commands := args.NewCollection()
	commands.Add(newMyCmd())
	commands.Add(newMyOtherCmd())

	// run the collection and handle errors
	exit, err := commands.Run()
	if err != nil {
		logger.Error(err.Error())
	}

	os.Exit(exit)
```

  
command implementation
```
type templateCmd struct {
}

// implement args.SubCmd interface

func (c *newMyCmd) ListArgs() []args.Arg {
	var a []args.Arg
	return a
}

func (c *newMyCmd) IsDefault() bool {
	return false
}

func (c *newMyCmd) Name() string {
	return "mycmd"
}

func (c *newMyCmd) Run() (int, error) {
	fmt.Println("My Command-1")
	return 0, nil
}

func (c *templateCmd) Help() string {
	return "help test for my command 1"
}

  ```
  command with arguments
```
type myOtherCmd struct {
	user     string
	password string
}

func newMyOtherCmd() *myOtherCmd {
	// default values for arguments 
	cmd := &myOtherCmd{
		user: "admin",
	}
	return cmd
}

func (cmd *myOtherCmd) ListArgs() []args.Arg {
	var a []args.Arg
	a = append(a, args.String(&cmd.user, "user", "User name for login").ShortCode('u'))
	a = append(a, args.String(&cmd.password, "password", "Password for user").ShortCode('p').Mandatory())
	return a
}

```
