package cmd

import (
	"log"

	"github.com/acygol/huntstat/framework"
)

//
// RegisterCommand is executed when someone calls 's!register'
// and registers a discord user to a community
//
func RegisterCommand(ctx framework.Context) {
	if !ctx.Cmd.ValidateArgs(len(ctx.Args)) {
		return
	}

	// args[0] must be a discord mention:
	if !framework.IsDiscordMention(ctx.Args[0]) {
		ctx.Reply("Invalid discord user. Are you mentioning them (@user)?")
		return
	}

	/*
		TODO: validate hunter profile, widget page returns 'invalid user' as body text
	*/

	_, err := framework.NewUser(ctx, framework.GetIDFromMention(ctx.Args[0]), ctx.Args[1])
	if err != nil {
		if err == framework.ErrAlreadyInGuild {
			ctx.Reply("User is already registered in this community.")
		} else {
			ctx.Reply("Failed to register user. Contact the bot maintainer for more information.")
		}
		log.Println("Failed to register user (", ctx.Guild.ID, ctx.Args[0], ctx.Args[1], "),", err)
		return
	}
	ctx.Reply("User " + ctx.Args[0] + " registered as " + ctx.Args[1])
}
