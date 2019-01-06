package framework

type (
	command func(Context)
	cmdMap  map[string]command

	//
	// CommandHandler is a struct holding the
	// collection of commands
	//
	CommandHandler struct {
		cmds cmdMap
	}
)

//
// NewCommandHandler initiates an instance of
// CommandHandler and returns a pointer to it
//
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(cmdMap)}
}

//
// GetCmds returns a map of all registered commands
// for a given CommandHandler
//
func (handler CommandHandler) GetCmds() cmdMap {
	return handler.cmds
}

//
// Get returns a pointer to the command function
// associated with the argument name
//
func (handler CommandHandler) Get(name string) (*command, bool) {
	cmd, found := handler.cmds[name]
	return &cmd, found
}

//
// Register adds a new command function and its stringified name
// to the given CommandHandler
//
func (handler CommandHandler) Register(name string, command command) {
	handler.cmds[name] = command
	if len(name) > 1 {
		handler.cmds[name[:1]] = command
	}
}
