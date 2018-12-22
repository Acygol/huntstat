package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
	"strings"
)

func DeleteCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!delete <@user>")
		return
	}
	if !strings.HasPrefix(ctx.Args[0], "<@") {
		ctx.Reply("Invalid user")
		return
	}
	if !framework.IsUserRegistered(ctx, ctx.Args[0]) {
		ctx.Reply("User is not registered")
		return
	}

	_, err := ctx.Conf.DbHandle.Exec("DELETE FROM users WHERE discord_name = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID)
	if err != nil {
		fmt.Println("error while executing query,", err)
		return
	}

	ctx.Reply("User removed")
}