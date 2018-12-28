package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
)

func DeleteCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!delete <@user>")
		return
	}

	if !framework.IsUserRegistered(ctx, ctx.Args[0]) {
		ctx.Reply("User is not registered")
		return
	}

	var err error
	if framework.IsDiscordMention(ctx.Args[0]) {
		_, err = ctx.Conf.Database.Handle.Exec("DELETE FROM users WHERE discord_id = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID)
	} else {
		_, err = ctx.Conf.Database.Handle.Exec("DELETE FROM users WHERE hunter_name = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID)
	}

	if err != nil {
		fmt.Println("error while executing query,", err)
		return
	}

	ctx.Reply("User removed")
}
