package cmd

import (
	"fmt"

	"github.com/acygol/huntstat/framework"
)

//
// DeleteCommand is executed when someone calls 's!unregister'
// and unregisters a user from a community
//
func DeleteCommand(ctx framework.Context) {
	if !framework.IsAdministrator(ctx.Discord, ctx.Guild, ctx.User) {
		ctx.Reply("You do not have permission to use this command.")
		return
	}

	if !ctx.Cmd.ValidateArgs(len(ctx.Args)) {
		ctx.Reply(fmt.Sprintf("Invalid syntax: s!%s %s", ctx.Cmd.Name, ctx.Cmd.CmdSyntax))
		return
	}

	user, err := framework.GetUserFromDiscordID(framework.GetIDFromMention(ctx.Args[0]))
	if err != nil {
		ctx.Reply("User is not registered to this guild.")
		return
	}
	err = user.RemoveFromGuild(ctx.Guild.ID, *ctx.Conf)
	if err != nil {
		ctx.Reply("Failed to remove user. Contact the bot maintainer for more information")
		fmt.Printf("Failed to remove user (%s, %s), %v", ctx.Guild.ID, ctx.Args[0], err)
		return
	}
	ctx.Reply("User removed")
}
