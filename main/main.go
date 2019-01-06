package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/acygol/huntstat/cmd"
	"github.com/acygol/huntstat/framework"
	"github.com/bwmarrin/discordgo"
)

var (
	config     *framework.Config
	cmdHandler *framework.CommandHandler
	botID      string
)

func main() {
	// load theHunter data
	framework.LoadGameData()

	// load config
	config = framework.NewConfig()

	// establish a command handler
	cmdHandler = framework.NewCommandHandler()
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
	botID = usr.ID

	// commandHandler is a callback for MessageCreate events
	// it ought to handle commands
	disc.AddHandler(commandHandler)
	disc.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		disc.UpdateStatus(0, "theHunter Classic")
		fmt.Println("HuntStat is running in", len(discord.State.Guilds), "guilds.")
	})
	//disc.AddHandler(framework.OnGuildJoined)

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
	if user.ID == botID || user.Bot {
		return
	}
	content := msg.Content

	// the message doesn't start with the bot's prefix
	if !strings.HasPrefix(content, config.Prefix) {
		return
	}

	// message only consists of the prefix, or not even that
	if len(content) <= len(config.Prefix) {
		return
	}

	// remove the prefix from the message's content
	content = content[len(config.Prefix):]
	if len(content) < 1 {
		return
	}

	// command arguments
	args := strings.Fields(content)
	name := strings.ToLower(args[0])

	command, found := cmdHandler.Get(name)
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
	ctx := framework.NewContext(sess, guild, channel, user, msg, config, cmdHandler)
	ctx.Args = args[1:]

	c := *command
	c(*ctx)
}

func registerCommands() {
	cmdHandler.Register("info", cmd.InfoCommand)
	cmdHandler.Register("help", cmd.InfoCommand)

	//
	// generate random hunt conditions
	//
	cmdHandler.Register("reserve", cmd.ReservesCommand)
	cmdHandler.Register("reserves", cmd.ReservesCommand)
	cmdHandler.Register("map", cmd.ReservesCommand)
	cmdHandler.Register("maps", cmd.ReservesCommand)

	cmdHandler.Register("weapon", cmd.WeaponsCommand)
	cmdHandler.Register("weapons", cmd.WeaponsCommand)
	cmdHandler.Register("gun", cmd.WeaponsCommand)
	cmdHandler.Register("guns", cmd.WeaponsCommand)

	cmdHandler.Register("animal", cmd.AnimalsCommand)
	cmdHandler.Register("animals", cmd.AnimalsCommand)

	cmdHandler.Register("modifier", cmd.ModifierCommand)
	cmdHandler.Register("modifiers", cmd.ModifierCommand)

	cmdHandler.Register("theme", cmd.ThemeCommand)
	cmdHandler.Register("themes", cmd.ThemeCommand)

	//
	// register process
	//
	cmdHandler.Register("register", cmd.RegisterCommand)
	cmdHandler.Register("unregister", cmd.DeleteCommand)
	cmdHandler.Register("delete", cmd.DeleteCommand)
	cmdHandler.Register("remove", cmd.DeleteCommand)

	//
	// generating widget links
	//
	cmdHandler.Register("widget", cmd.WidgetCommand)

	//
	// leaderboard
	//
	cmdHandler.Register("leaderboard", cmd.LeaderboardCommand)
	cmdHandler.Register("leaderboards", cmd.LeaderboardCommand)
}
