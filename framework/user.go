package framework

import (
	"log"
)

func NewUser(ctx Context, guildid, discordname, huntername string) error {
	stmt, err := ctx.Conf.DbHandle.Prepare("INSERT INTO users(guild_id, discord_name, hunter_name) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(guildid, discordname, huntername)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func IsUserRegistered(ctx Context, discordname string) bool {
	var rows int
	ctx.Conf.DbHandle.QueryRow("SELECT COUNT(*) FROM users WHERE discord_name = ? AND guild_id = ?", discordname, ctx.Guild.ID).Scan(&rows)
	return rows > 0
}
