package framework

import (
	"log"
)

func NewUser(ctx Context, guildid string, discordname string, huntername string) error {
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
