package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
)

func RegisterCommand(ctx framework.Context) {
	if len(ctx.Args) < 2 {
		ctx.Reply("Invalid syntax: s!register <@discord_name | name#0000> <hunter_name>")
		return
	}

	// fmt.Printf("[ctx: %v], [guildid: %s], [args(0): %s], [args(1): %s]", ctx, ctx.Guild.ID, ctx.Args[0], ctx.Args[1])

	/*
		TODO: validate arguments
	*/
	err := framework.NewUser(ctx, ctx.Guild.ID, ctx.Args[0], ctx.Args[1])
	if err != nil {
		ctx.Reply("Failed to register new user. Contact the bot maintainer for more information")
		fmt.Println("Failed to register new user (", ctx.Guild.ID, ctx.Args[0], ctx.Args[1], "),", err)
		return
	}
	ctx.Reply("User " + ctx.Args[0] + " registered as " + ctx.Args[1])
}
