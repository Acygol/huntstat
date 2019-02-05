package cmd

import (
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
	if !ctx.CmdHandler.MustGet(ctx.CmdName).ValidateArgs(ctx) {
		return
	}
	processProfileQuery(ctx, UhcURL)
}

//
// ProfileCommand is executed when someone calls 's!profile'
//
func ProfileCommand(ctx framework.Context) {
	if !ctx.CmdHandler.MustGet(ctx.CmdName).ValidateArgs(ctx) {
		return
	}
	processProfileQuery(ctx, ProfileURL)
}

//
// WidgetCommand is executed when someone calls 's!widget'
//
func WidgetCommand(ctx framework.Context) {
	if !ctx.CmdHandler.MustGet(ctx.CmdName).ValidateArgs(ctx) {
		return
	}
	processProfileQuery(ctx, WidgetURL)
}

func processProfileQuery(ctx framework.Context, whichprofile int) {
	var reply strings.Builder

	if framework.IsDiscordMention(ctx.Args[0]) {
		user, err := framework.GetUserFromDiscordID(framework.GetIDFromMention(ctx.Args[0]))
		if err != nil || !user.IsInGuild(ctx.Guild.ID) {
			ctx.Reply("User is not registered")
			return
		}
		fmt.Fprintf(&reply, "%s for %s\n", getURLHeader(whichprofile), ctx.Args[0])
		fmt.Fprintf(&reply, "<%s>\n", GetURL(user.HunterName(), whichprofile))
	} else if strings.EqualFold(ctx.Args[0], "all") {
		users := framework.GetUsersInGuild(ctx.Guild.ID)
		for _, user := range users {
			fmt.Fprintf(&reply, "<%s>\n", GetURL(user.HunterName(), whichprofile))
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
		return fmt.Sprintf("https://www.uhcapps.co.uk/stats.php?username=%s", huntername)
	case ProfileURL:
		return fmt.Sprintf("https://www.thehunter.com/#profile/%s/", huntername)
	}
	return fmt.Sprintf("http://widget.thehunter.com/signature/?user=%s", huntername)
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
