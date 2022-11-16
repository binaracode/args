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

// implement args.SubCmd.ListArgs
// returns a list of arguments for this command (no arguments for this example)
func (c *newMyCmd) ListArgs() []args.Arg {
	var a []args.Arg
	return a
}

// implement args.SubCmd.IsDefault
// returns true if this command is the default command, in case if the application was
// started with no sub command
func (c *newMyCmd) IsDefault() bool {
	return false
}

// implement args.SubCmd.Name
// returns the name for the sub command
func (c *newMyCmd) Name() string {
	return "mycmd"
}

// implement args.SubCmd.ListArgs
// the body of the sub command goes here, return the application exit value and/or
// an error
func (c *newMyCmd) Run() (int, error) {
	fmt.Println("My Command-1")
	return 0, nil
}

// implement args.SubCmd.Help
// returns the usage information of the sub command, argument usage/help
// is added with the arguments in ListArgs function
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

// implement args.SubCmd.ListArgs
// returns a list of arguments for this command
// user argument has a short version (-u or /u), if not passed the default user 'admin' is used
// as initiaalized in newMyOtherCmd function and password is a mandatory argument
func (cmd *myOtherCmd) ListArgs() []args.Arg {
	var a []args.Arg
	a = append(a, args.String(&cmd.user, "user", "User name for login").ShortCode('u'))
	a = append(a, args.String(&cmd.password, "password", "Password for user").ShortCode('p').Mandatory())
	return a
}

```
