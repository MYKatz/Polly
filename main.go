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
	discord, err := discordgo.New("Bot " + botKey)
	checkErr("Error creating discord session", err)
	user, err := discord.User("@me")
	checkErr("Error retrieving bot account", err)

	botID = user.ID

}

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %+v", msg, err))
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