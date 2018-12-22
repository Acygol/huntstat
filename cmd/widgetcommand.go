package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
	"strings"
)

func WidgetCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Missing argument: s!widget <all | @user>")
		return
	}
	// result value is stored in this variable
	var huntername string

	// retrieve this user's widget
	if strings.HasPrefix(ctx.Args[0], "<@") {
		err := ctx.Conf.DbHandle.QueryRow("SELECT hunter_name FROM users WHERE discord_name = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID).Scan(&huntername)
		if err != nil {
			ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information")
			fmt.Println("error retrieving from database,", err)
			return
		}
		ctx.Reply("Widget for " + ctx.Args[0])
		ctx.Reply("http://widget.thehunter.com/signature/?user=" + huntername)
	}

	// go through all users to retrieve their hunter name
	if strings.EqualFold(ctx.Args[0], "all") {
		rows, err := ctx.Conf.DbHandle.Query("SELECT hunter_name FROM users WHERE guild_id = ?", ctx.Guild.ID)
		if err != nil {
			ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information")
			fmt.Println("error retrieving from database,", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&huntername)
			if err != nil {
				fmt.Println("Error attempting to scan the next row,", err)
				break
			}
			ctx.Reply("http://widget.thehunter.com/signature/?user=" + huntername)
		}
	}
}
