package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/acygol/huntstat/framework"
)

const (
	//
	// UhcURL is an iota constant to tell various functions
	// such as: GetURL() which profile URL to return. When
	// used, GetURL() and functions alike will return the data
	// related to the UHCApps website
	//
	UhcURL = iota

	//
	// ProfileURL is an iota constant to tell various functions
	// such as: GetURL() which profile URL to return. When
	// used, GetURL() and functions alike will return the data
	// related to the TheHunter website
	//
	ProfileURL = iota

	//
	// WidgetURL is an iota constant to tell various functions
	// such as: GetURL() which profile URL to return. When
	// used, GetURL() and functions alike will return the data
	// related to the TheHunter widget
	//
	WidgetURL = iota
)

const (
	uhcHeader     = "UHCApps URL"
	profileHeader = "TheHunter profile"
	widgetHeader  = "Widget"
)

//
// UhcCommand is executed when someone calls 's!uhc'
//
func UhcCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!uhc <@user | all>")
		return
	}
	processProfileQuery(ctx, UhcURL)
}

//
// ProfileCommand is executed when someone calls 's!profile'
//
func ProfileCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!profile <@user | all>")
		return
	}
	processProfileQuery(ctx, ProfileURL)
}

//
// WidgetCommand is executed when someone calls 's!widget'
//
func WidgetCommand(ctx framework.Context) {
	if len(ctx.Args) < 1 {
		ctx.Reply("Invalid syntax: s!widget <@user | all>")
		return
	}
	processProfileQuery(ctx, WidgetURL)
}

func processProfileQuery(ctx framework.Context, whichprofile int) {
	var huntername string
	var reply strings.Builder
	if framework.IsDiscordMention(ctx.Args[0]) {
		// Mentions that start with '<@' are valid server members
		err := ctx.Conf.Database.Handle.QueryRow("SELECT hunter_name FROM users WHERE discord_id = ? AND guild_id = ?", ctx.Args[0], ctx.Guild.ID).Scan(&huntername)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.Reply("User isn't registered")
			} else {
				ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information")
			}
			fmt.Println("error retrieving from database,", err)
			return
		}
		fmt.Fprintf(&reply, "%s for %s\n", getURLHeader(whichprofile), ctx.Args[0])
		fmt.Fprintf(&reply, GetURL(huntername, whichprofile))
	} else if strings.EqualFold(ctx.Args[0], "all") {
		rows, err := ctx.Conf.Database.Handle.Query("SELECT hunter_name FROM users WHERE guild_id = ?", ctx.Guild.ID)
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
			fmt.Fprintf(&reply, "%s\n", GetURL(huntername, whichprofile))
		}
	} else {
		fmt.Fprintln(&reply, "Invalid user")
	}
	ctx.Reply(reply.String())
}

//
// GetURL returns a non-embedding URL to the provided hunter's
// profile for a specific website depending on the value of
// whichprofile
//
func GetURL(huntername string, whichprofile int) string {
	switch whichprofile {
	case UhcURL:
		return fmt.Sprintf("<https://www.uhcapps.co.uk/stats.php?username=%s>", huntername)
	case ProfileURL:
		return fmt.Sprintf("<https://www.thehunter.com/#profile/%s/>", huntername)
	}
	return fmt.Sprintf("<http://widget.thehunter.com/signature/?user=%s>", huntername)
}

func getURLHeader(whichprofile int) string {
	switch whichprofile {
	case UhcURL:
		return fmt.Sprint(uhcHeader)
	case ProfileURL:
		return fmt.Sprint(profileHeader)
	}
	return fmt.Sprint(widgetHeader)
}
