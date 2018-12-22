package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
)

func RegisterCommand(ctx framework.Context) {
	if len(ctx.Args) < 2 {
		ctx.Reply("Invalid syntax: s!register <@user> <hunter_name>")
		return
	}

	// it's an invalid mention
	if !framework.IsDiscordMention(ctx.Args[0]) {
		ctx.Reply("Invalid user")
		return
	}

	/*
		TODO: validate hunter profile, widget page returns 'invalid user' as body text
	*/

	// adds a new user to the database
	err := framework.NewUser(ctx, ctx.Guild.ID, ctx.Args[0], ctx.Args[1])
	if err != nil {
		ctx.Reply("Failed to register new user. Contact the bot maintainer for more information")
		fmt.Println("Failed to register new user (", ctx.Guild.ID, ctx.Args[0], ctx.Args[1], "),", err)
		return
	}
	ctx.Reply("User " + ctx.Args[0] + " registered as " + ctx.Args[1])
}
