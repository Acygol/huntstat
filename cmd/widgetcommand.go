package cmd

import (
	"database/sql"
	"fmt"
	"github.com/acygol/huntstat/framework"
	"strings"
)

func WidgetCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!widget <@user | all>")
		return
	}

	// result value is stored in this variable
	var huntername string
	var reply strings.Builder

	if framework.IsDiscordMention(ctx.Args[0]) {
		// Mentions that start with '<@' are valid server members
		err := ctx.Conf.DbHandle.QueryRow("SELECT hunter_name FROM users WHERE discord_name = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID).Scan(&huntername)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.Reply("User isn't registered")
			} else {
				ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information")
			}
			fmt.Println("error retrieving from database,", err)
			return
		}
		fmt.Fprintf(&reply, "Widget for %s\n", ctx.Args[0])
		fmt.Fprintf(&reply, GetWidget(huntername))
	} else if strings.EqualFold(ctx.Args[0], "all") {
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
			fmt.Fprintf(&reply, GetWidget(huntername))
		}
	} else {
		fmt.Fprintln(&reply, "Invalid user")
	}
	ctx.Reply(reply.String())
}

func GetWidget(huntername string) string {
	var url strings.Builder
	fmt.Fprintf(&url, "http://widget.thehunter.com/signature/?user=%s", huntername)
	return url.String()
}
