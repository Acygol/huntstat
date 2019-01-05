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
	role, err := validateRole(sess, event.Guild, guildRoles, "user", true)
	if err != nil {
		log.Printf("")
		return
	}

	// give everyone the user role
	for _, member := range event.Guild.Members {
		assignRoleToUser(sess, event.Guild, member, role)
	}

	validateRole(sess, event.Guild, guildRoles, "admin", false)
}

func validateRole(sess *discordgo.Session, guild *discordgo.Guild, roles []*discordgo.Role, name string, hoist bool) (*discordgo.Role, error) {
	exists := doesRoleExist(guild, name)
	var role *discordgo.Role
	var err error
	if !exists {
		log.Printf("role '%s' does not exist in guild (%s), creating...", name, guild.Name)
		role, err = createRole(sess, guild, name, hoist)
		if err != nil {
			log.Printf("error creating role,", err)
			return nil, err
		}
	} else {
		role = roles[getRoleID(guild, name)]
	}
	return role, nil
}

func doesRoleExist(guild *discordgo.Guild, name string) bool {
	return (getRoleID(guild, name) != -1)
}

func getRoleID(guild *discordgo.Guild, name string) int {
	for i, role := range guild.Roles {
		if role.Name == name {
			return i
		}
	}
	return -1
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
