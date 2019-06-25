package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	commandPrefix string
	botID string
	botKey string
)

func main() {
	commandPrefix, botKey := getConfigVars()
	fmt.Printf("Initializing Polly with command prefix '%s' \n", commandPrefix)
	discord, err := discordgo.New("Bot " + botKey)
	checkErr("Error creating discord session", err)
	user, err := discord.User("@me")
	checkErr("Error retrieving bot account", err)

	botID = user.ID
	//handlers. There are many different types in the library, corresponding to each of these event types https://discordapp.com/developers/docs/topics/gateway#event-names
	discord.AddHandler(commandHandler)
	discord.AddHandler(readyHandler)
	err = discord.Open()
	checkErr("Unable to open a connection to discord: ", err)

	defer discord.Close()

	//incoming hacky thing - this creates a channel, preventing our main() function from closing on its own, so the bot stays alive while
	<-make(chan struct{})
}

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %+v", msg, err))
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		return
	}
	content := message.Content
	if string(content[0]) != commandPrefix{
		fmt.Printf("message received, but doesn't start with !: %s \n", content)
		return
	}
	//do something with the content here later!
	//for now, we'll just print it
	fmt.Printf("From %s: '%s'", message.Author, content)
	return
}

func readyHandler(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "Polly want a cracker")
	if err != nil {
		panic(fmt.Errorf("Fatal error, could not update status: %s", err))
	}
	servers := discord.State.Guilds //returns an array of all servers the bot is added to
	fmt.Printf("I'm installed on %d servers. Nice! \n", len(servers))
}

func disconnect(discord *discordgo.Session) {
	//set status to idle (-1)
	err := discord.UpdateStatus(-1, "")
	if err != nil {
		panic(fmt.Errorf("Fatal error, could not update status: %s", err))
	}
}

func getConfigVars() (string, string) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error, check config file/environment variables: %s \n", err))
	}
	prefix := viper.GetString("COMMAND_PREFIX")
	if prefix == "" || len(prefix) != 1 {
		panic(fmt.Errorf("Fatal error, check COMMAND_PREFIX environment variable"))
	}
	key := viper.GetString("BOT_KEY")
	if key == "" {
		panic(fmt.Errorf("Fatal error, check BOT_KEY environment variable"))
	}
	return prefix, key
}