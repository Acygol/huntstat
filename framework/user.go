package framework

import (
	"errors"
	"log"
)

type (
	//
	// User ...
	//
	User struct {
		databaseID int64
		discordID  string
		hunterName string
		guilds     []string
	}
)

var (
	//
	// Users holds all registered users
	//
	Users []*User

	//
	// ErrUserNotFound is returned by GetUserFromDiscordID when
	// the user is not loaded in. This is a sign that the requested
	// user is not registered to the system yet
	//
	ErrUserNotFound = errors.New("user: no user found")

	//
	// ErrAlreadyInGuild is returned by User.RegisterGuild when
	// the provided guild is already in the user's 'guilds' slice
	//
	ErrAlreadyInGuild = errors.New("user: already in guild")

	//
	// ErrNotInGuild is returned by User.RemoveFromGuild when
	// the provided guild is not in the user's 'guilds' slice
	//
	ErrNotInGuild = errors.New("user: not in guild")
)

//
// NewUser ...
//
func NewUser(ctx Context, discordID, hunterName string) (User, error) {
	user, err := GetUserFromDiscordID(discordID)
	if err != nil {
		// user isn't registered yet
		stmt, err := ctx.Conf.Database.Handle.Prepare("insert into user (discord_id, hunter_name) values (?, ?)")
		if err != nil {
			log.Println("error while preparing NewUser query,", err)
			return User{}, err
		}
		res, err := stmt.Exec(discordID, hunterName)
		if err != nil {
			log.Println("error while executing NewUser query,", err)
			return User{}, err
		}
		lastid, _ := res.LastInsertId()
		user = &User{
			databaseID: lastid,
			discordID:  discordID,
			hunterName: hunterName,
			guilds:     []string{},
		}
		Users = append(Users, user)
	}
	err = user.RegisterGuild(ctx.Guild.ID, *ctx.Conf)
	return *user, err
}

//
// LoadUsers queries the database and loads all the records
// in memory.
//
func LoadUsers(conf Config) {
	rows, err := conf.Database.Handle.Query("select user.*, user_guilds.guild_id from user inner join user_guilds on user.id = user_guilds.user_id")
	if err != nil {
		log.Fatal("error while querying database,", err)
		return
	}
	defer rows.Close()

	var (
		databaseID int64
		discordID  string
		hunterName string
		guildID    string
	)

	for highestID := int64(0); rows.Next(); {
		err := rows.Scan(&databaseID, &discordID, &hunterName, &guildID)
		if err != nil {
			log.Fatal("error scanning row,", err)
			return
		}

		//
		// highestID == databaseID when a user is part of multiple guilds
		// that use HuntStat. If so, the value of 'guildID' is appended
		// to the user's 'guilds' slice
		//
		if highestID == databaseID {
			user, err := GetUserFromDiscordID(discordID)
			if err != nil {
				log.Fatal("error while retrieving user from discord ID,", err)
				return
			}
			user.guilds = append(user.guilds, guildID)
		} else {
			tmpUser := &User{
				databaseID: databaseID,
				discordID:  discordID,
				hunterName: hunterName,
				guilds:     []string{guildID},
			}
			Users = append(Users, tmpUser)
		}
		highestID = databaseID
	}
}

//
// GetUserFromDiscordID ...
//
func GetUserFromDiscordID(discordID string) (*User, error) {
	for _, user := range Users {
		if user.discordID == discordID {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

//
// HunterName is a getter method
//
func (user User) HunterName() string {
	return user.hunterName
}

//
// GetGuildIndex ...
//
func (user User) GetGuildIndex(guildID string) int {
	for index, gID := range user.guilds {
		log.Println(guildID)
		if gID == guildID {
			return index
		}
	}
	return -1
}

//
// IsInGuild ...
//
func (user User) IsInGuild(guildID string) bool {
	return (user.GetGuildIndex(guildID) != -1)
}

//
// RegisterGuild ...
//
func (user *User) RegisterGuild(guildID string, conf Config) error {
	if user.IsInGuild(guildID) {
		return ErrAlreadyInGuild
	}
	stmt, err := conf.Database.Handle.Prepare("insert into user_guilds (user_id, guild_id) values (?, ?)")
	if err != nil {
		log.Println("error while preparing RegisterGuild query,", err)
		return err
	}
	_, err = stmt.Exec(user.databaseID, guildID)
	if err != nil {
		log.Println("error while executing RegisterGuild query,", err)
		return err
	}
	user.guilds = append(user.guilds, guildID)

	// DEBUG
	for _, gID := range user.guilds {
		log.Println("RegiserGuild: ", gID)
	}
	return nil
}

//
// RemoveFromGuild ...
//
func (user *User) RemoveFromGuild(guildID string, conf Config) error {
	index := user.GetGuildIndex(guildID)
	if index == -1 {
		return ErrNotInGuild
	}
	stmt, err := conf.Database.Handle.Prepare("delete from user_guilds where user_id = ? and guild_id = ?")
	if err != nil {
		log.Println("error while preparing RemoveFromGuild query,", err)
		return err
	}
	_, err = stmt.Exec(user.databaseID, guildID)
	if err != nil {
		log.Println("error while executing RemoveFromGuild query,", err)
		return err
	}
	user.guilds = append(user.guilds[:index], user.guilds[index+1:]...)

	if len(user.guilds) < 1 {
		_, err := conf.Database.Handle.Exec("delete from user where id = ?", user.databaseID)
		if err != nil {
			log.Println("error while removing user from database,", err)
			return err
		}
		for i, u := range Users {
			if u.databaseID == user.databaseID {
				Users = append(Users[:i], Users[i+1:]...)
				break
			}
		}
	}
	return nil
}

//
// GetUsersInGuild returns a slice of users that are registered to
// the provided guild ID
//
func GetUsersInGuild(guildID string) (users []*User) {
	for _, user := range Users {
		for _, gID := range user.guilds {
			if gID == guildID {
				users = append(users, user)
			}
		}
	}
	return users
}
