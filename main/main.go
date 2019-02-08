package main

import (
	"fmt"
	"log"
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
	if config == nil {
		log.Fatal("error loading config file")
		return
	}

	// load users
	framework.LoadUsers(*config)

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

	// commandHandler is a callback for MessageCreate events to handle commands
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
	ctx := framework.NewContext(sess, guild, channel, user, msg, config, cmdHandler, command, args[1:])

	c := command.CmdFunc
	c(*ctx)
}

func registerCommands() {
	command := cmdHandler.Register("info", cmd.InfoCommand)
	command.Description("Prints all commands that the bot understands")
	command.Syntax("<(optional) command name>")
	command.RegisterAlias("help")

	//
	// generate random hunt conditions
	//
	command = cmdHandler.Register("reserve", cmd.ReservesCommand)
	command.Description("Generates a random reserve to hunt on")
	command.RegisterAlias("reserves")
	command.RegisterAlias("map")
	command.RegisterAlias("maps")

	command = cmdHandler.Register("weapon", cmd.WeaponsCommand)
	command.Description("Generates a random weapon loadout to hunt with")
	command.Syntax("<reserve> <max inventory capacity>")
	command.RegisterAlias("weapons")
	command.RegisterAlias("gun")
	command.RegisterAlias("guns")

	command = cmdHandler.Register("animal", cmd.AnimalsCommand)
	command.Description("Generates a random animal to hunt on")
	command.Syntax("<(optional) reserve>")
	command.RegisterAlias("animals")

	command = cmdHandler.Register("modifier", cmd.ModifierCommand)
	command.Description("Generates a random modifier to your hunt")
	command.RegisterAlias("modifiers")

	command = cmdHandler.Register("theme", cmd.ThemeCommand)
	command.Description("Generates a random hunt theme")
	command.RegisterAlias("themes")

	//
	// register process
	//
	command = cmdHandler.Register("register", cmd.RegisterCommand)
	command.Description("registers a user to the community")
	command.Syntax("<@user> <hunter name>")
	command.RegisterAlias("add")

	command = cmdHandler.Register("unregister", cmd.DeleteCommand)
	command.Description("removes a user from the community")
	command.Syntax("<@user>")
	command.RegisterAlias("delete")
	command.RegisterAlias("remove")

	//
	// generating profile links
	//
	command = cmdHandler.Register("widget", cmd.WidgetCommand)
	command.Description("generates the widget URL for a given community member")
	command.Syntax("<@user | all>")

	command = cmdHandler.Register("profile", cmd.ProfileCommand)
	command.Description("generates the profile URL for a given community member")
	command.Syntax("<@user | all>")

	command = cmdHandler.Register("uhc", cmd.UhcCommand)
	command.Description("generates the UHC statistics URL for a given community member")
	command.Syntax("<@user | all>")

	//
	// leaderboard
	//
	command = cmdHandler.Register("leaderboard", cmd.LeaderboardCommand)
	command.Description("generates a leaderboard for the community")
	command.Syntax("<(optional) animal name>")
	command.RegisterAlias("leaderboards")
}
