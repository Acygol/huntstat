package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/acygol/huntstat/framework"
)

//
// InfoCommand is executed when someone calls 's!info'
//
func InfoCommand(ctx framework.Context) {
	if len(ctx.Args) > 0 {
		//
		// command name was given
		//
		command, found := ctx.CmdHandler.Get(ctx.Args[0])
		if !found {
			ctx.Reply("Invalid command")
		} else {
			var reply strings.Builder

			fmt.Fprintf(&reply, "s!%s: %s", ctx.Args[0], command.CmdDesc)
			if len(command.CmdSyntax) > 0 {
				fmt.Fprintf(&reply, "\nUsage: s!%s %s", ctx.Args[0], command.CmdSyntax)
			}
			ctx.Reply(reply.String())
		}
	} else {
		cmds := ctx.CmdHandler.GetCmds()
		buffer := bytes.NewBufferString("Commands: ")
		for key := range cmds {
			if len(key) == 1 {
				continue
			}
			buffer.WriteString(key)
			buffer.WriteString(", ")
		}
		str := buffer.String()
		ctx.Reply(str[:len(str)-2])
	}
}
