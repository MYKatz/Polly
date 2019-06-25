package main

import (
	"fmt"
	//"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	commandPrefix string
	botID string
	botKey string
)

func main() {
	commandPrefix, botKey := getConfigVars()
	fmt.Printf(commandPrefix)
	fmt.Printf(botKey)
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