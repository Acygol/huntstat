package framework

import (
	"fmt"
	"strings"
)

type (
	//
	// Command is a struct that holds
	// information related to a given
	// command
	//
	Command struct {
		CmdFunc   func(Context)
		CmdDesc   string
		CmdSyntax string
		aliases   []string
	}

	//
	// CmdMap is an associative array that
	// binds command names to their related
	// instance of command
	//
	CmdMap map[string]*Command

	//
	// CommandHandler is a struct holding the
	// collection of commands
	//
	CommandHandler struct {
		cmds CmdMap
	}
)

//
// NewCommandHandler initiates an instance of
// CommandHandler and returns a pointer to it
//
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

//
// GetCmds returns a map of all registered commands
// for a given CommandHandler
//
func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

//
// Get returns a pointer to the command function
// associated with the argument name
//
func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	//
	// if the map look-up fails, then the argument name
	// is probably an alias for a command
	//
	if !found {
		//
		// Loop over all commands registered to the handler
		//
		for _, cmd := range handler.cmds {
			//
			// Does the command have any aliases?
			//
			if len(cmd.aliases) > 0 {
				for _, alias := range cmd.aliases {
					if strings.EqualFold(alias, name) {
						return cmd, true
					}
				}
			}
		}
	}
	return cmd, found
}

//
// MustGet is a wrapper around CommandHandler.Get() that panics
// when it cannot find the given command. This method should only
// be used when you are certain that the command exists. When
// used properly, it allows for method chaining
//
func (handler CommandHandler) MustGet(name string) *Command {
	cmd, found := handler.Get(name)
	if !found {
		panic(`CommandHandler: MustGet(` + name + `): command not found`)
	}
	return cmd
}

//
// Register adds a new command function and its stringified name
// to the given CommandHandler.
//
func (handler CommandHandler) Register(name string, cmdFunc func(Context)) *Command {
	cmd := new(Command)
	cmd.CmdFunc = cmdFunc
	cmd.CmdSyntax = ""
	cmd.CmdDesc = "<not defined>"

	handler.cmds[name] = cmd
	if len(name) > 1 {
		handler.cmds[name[:1]] = cmd
	}
	return cmd
}

//
// Description adds a descriptive help message to
// the command and its aliases
//
func (command *Command) Description(desc string) {
	command.CmdDesc = desc
}

//
// RegisterAlias adds an alias name to the command
//
func (command *Command) RegisterAlias(alias string) {
	command.aliases = append(command.aliases, alias)
}

//
// Syntax defines how the given command must be used
//
func (command *Command) Syntax(syntax string) {
	command.CmdSyntax = syntax
}

//
// GetArgsCount returns the number of arguments the command
// expects
//
func (command Command) GetArgsCount() int {
	var args []string

	// exclude optional arguments
	for _, arg := range strings.Split(command.CmdSyntax, "<") {
		if !strings.Contains(arg, "optional") && len(arg) > 0 {
			args = append(args, arg)
		}
	}
	return len(args)
}

/*
func (command Command) HasOnlyOptionalArgs() bool {
	for _, arg := range strings.Split(command.CmdSyntax, "<") {
		if len(arg) < 1 {
			continue
		}
		if !strings.Contains(arg, "optional") {
			return false
		}
	}
	return true
}
*/

//
// ValidateArgs is a helper method to validate arguments
// of a given command
//
func (command Command) ValidateArgs(ctx Context) bool {
	if len(ctx.Args) < command.GetArgsCount() {
		ctx.Reply(fmt.Sprintf("Invalid syntax: s!%s %s", ctx.CmdName, command.CmdSyntax))
		return false
	}
	return true
}
