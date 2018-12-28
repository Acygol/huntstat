package framework

import (
	"log"
	"strings"
)

func NewUser(ctx Context, guildid, discordname, huntername string) error {
	stmt, err := ctx.Conf.Database.Handle.Prepare("INSERT INTO users(guild_id, discord_id, hunter_name) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(guildid, discordname, huntername)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func IsDiscordMention(discordname string) bool {
	return strings.HasPrefix(discordname, "<@")
}

//
// username can be either a hunter profile or a discord mention
func IsUserRegistered(ctx Context, username string) bool {
	var rows int
	if IsDiscordMention(username) {
		ctx.Conf.Database.Handle.QueryRow("SELECT COUNT(*) FROM users WHERE discord_id = ? AND guild_id = ?", username, ctx.Guild.ID).Scan(&rows)
	} else {
		ctx.Conf.Database.Handle.QueryRow("SELECT COUNT(*) FROM users WHERE hunter_name = ? AND guild_id = ?", username, ctx.Guild.ID).Scan(&rows)
	}
	return rows > 0
}
