package framework

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

//
// OnGuildJoined is called when the bot connects to a server.
// In this particular instance, it is used to verify the roles
// that the bot needs to function correctly.
//
// HuntStat uses 2 roles to determine command execution
// permissions:
//		1. admin — can execute all commands
//		2. user — can't execute commands like 'DeleteCommand'
//
func OnGuildJoined(sess *discordgo.Session, event *discordgo.GuildCreate) {
	fmt.Printf("Connected to %s\n", event.Guild.Name)

	guildRoles, err := sess.GuildRoles(event.Guild.ID)
	if err != nil {
		log.Println("error retrieving guild roles,", err)
		return
	}

	exists := doesRoleExist(guildRoles, "user")
	if !exists {
		log.Printf("role 'user' doesn't exist in guild (%s)", event.Guild.Name)

		// create user role
		/*
			role, err := createRole(sess, event.Guild, "user", true)
			if err != nil {
				log.Println("error creating role 'user',", err)
				return
			}
			// give everyone the user role
			for _, member := range event.Guild.Members {
				assignRoleToUser(sess, event.Guild, member, role)
			}
		*/
	}

	if exists = doesRoleExist(guildRoles, "admin"); !exists {
		log.Printf("role 'admin' doesn't exist in guild (%s)", event.Guild.Name)
		// create admin role
		/*
			_, err := createRole(sess, event.Guild, "admin", false)
			if err != nil {
				log.Println("error creating role 'admin',", err)
				return
			}
		*/
	}
}

func doesRoleExist(roles []*discordgo.Role, name string) bool {
	for _, role := range roles {
		if role.Name == name {
			return true
		}
	}
	return false
}

func createRole(sess *discordgo.Session, guild *discordgo.Guild, name string, hoist bool) (*discordgo.Role, error) {
	role, err := sess.GuildRoleCreate(guild.ID)
	if err != nil {
		return nil, err
	}

	role, err = sess.GuildRoleEdit(
		guild.ID,
		role.ID,
		role.Name,
		role.Color,
		hoist,
		role.Permissions,
		false,
	)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func assignRoleToUser(sess *discordgo.Session, guild *discordgo.Guild, member *discordgo.Member, role *discordgo.Role) {
	err := sess.GuildMemberRoleAdd(guild.ID, member.User.ID, role.ID)
	if err != nil {
		log.Println("error assigning role to user,", err)
		return
	}
}
