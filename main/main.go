package main

import (
	"fmt"
	"github.com/acygol/huntstat/cmd"
	"github.com/acygol/huntstat/framework"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	config 		*framework.Config
	CmdHandler 	*framework.CommandHandler
	botId		string
)

func main() {
	// load config
	config = framework.Init("../config.json")
	if config == nil {
		fmt.Println("error initializing config")
		return
	}

	// establish a command handler
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	// establish discord session
	disc, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	usr, err := disc.User("@me")
	if err != nil {
		fmt.Println("Error obtaining user details,", err)
		return
	}
	botId = usr.ID

	// commandHandler is a callback for MessageCreate events
	// it ought to handle commands
	disc.AddHandler(commandHandler)
	disc.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		disc.UpdateStatus(0, "theHunter Classic")
		guilds := disc.State.Guilds
		fmt.Println("HuntStat is running in", len(guilds), "guilds.")
	})

	// Open a websocket connection to Discord
	err = disc.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait for process interruption signal (i.e., CTRL + C)
	fmt.Println("HuntStat is running.	Press CTRL+C to stop.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Closes the discord session
	disc.Close()
}

func commandHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author

	// ignore the message when the bot themselves sent it
	if user.ID == botId || user.Bot {
		return
	}
	content := msg.Content

	if !strings.HasPrefix(content, config.Prefix) {
		return
	}

	// message only consists of the prefix, or not even that
	if len(content) <= len(config.Prefix) {
		return
	}

	// content without prefix
	content = content[len(config.Prefix):]
	if len(content) < 1 {
		return
	}

	// command arguments
	args := strings.Fields(content)
	name := strings.ToLower(args[0])

	command, found := CmdHandler.Get(name)
	if !found {
		sess.ChannelMessageSend(msg.ChannelID, "Command not found")
		return
	}
	channel, err := sess.State.Channel(msg.ChannelID)
	if err != nil {
		fmt.Println("Error while retrieving channel,", err)
		return
	}
	guild, err := sess.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error while retrieving guild,", err)
		return
	}
	ctx := framework.NewContext(sess, guild, channel, user, msg, config, CmdHandler)
	ctx.Args = args[1:]

	c := *command
	c(*ctx)
}

func registerCommands() {
	//CmdHandler.Register("info", cmd.InfoCommand)
	//CmdHandler.Register("help", cmd.InfoCommand)

	/*
		generate random hunt conditions
	*/
	CmdHandler.Register("reserves", cmd.ReservesCommand)
	CmdHandler.Register("reserve", cmd.ReservesCommand)
	CmdHandler.Register("maps", cmd.ReservesCommand)
	CmdHandler.Register("map", cmd.ReservesCommand)

	CmdHandler.Register("weapons", cmd.WeaponsCommand)
	CmdHandler.Register("weapon", cmd.WeaponsCommand)
	CmdHandler.Register("gun", cmd.WeaponsCommand)
	CmdHandler.Register("guns", cmd.WeaponsCommand)

	CmdHandler.Register("animals", cmd.AnimalsCommand)
	CmdHandler.Register("animal", cmd.AnimalsCommand)

	CmdHandler.Register("modifier", cmd.ModifierCommand)
	CmdHandler.Register("modifiers", cmd.ModifierCommand)

	CmdHandler.Register("themes", cmd.ThemeCommand)
	CmdHandler.Register("theme", cmd.ThemeCommand)

	// register users
	CmdHandler.Register("register", cmd.RegisterCommand)

	// generating widget links
	CmdHandler.Register("widget", cmd.WidgetCommand)
}
