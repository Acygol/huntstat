package framework

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

/*
// Context is a struct holding all relevant information
// regarding a discord guild when a command is executed
// It allows for commands to recognize in which guild
// and which channel the command was issued in, and
// by who.
// It is passed as a function argument to command
// functions. Could also be used in non-command
// calls if needed.
*/
type Context struct {
	Discord			*discordgo.Session
	Guild			*discordgo.Guild
	TextChannel 	*discordgo.Channel
	User			*discordgo.User
	Message			*discordgo.MessageCreate
	Args 			[]string

	Conf			*Config
	CmdHandler		*CommandHandler
}

/*
// NewContext populates an instance of Context, 
// returning it as a pointer
*/
func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel, user *discordgo.User,
				message *discordgo.MessageCreate, conf *Config, cmdHandler *CommandHandler) *Context {

	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.Conf = conf
	ctx.CmdHandler = cmdHandler
	return ctx
}

/*
// Reply acts as a wrapper for 
// discordgo.Session.ChannelMessageSend
*/
func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}
