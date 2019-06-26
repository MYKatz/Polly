//Various commands

package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func markdownWrapper(lang string, message string) string {
	//hack for working with backticks in go strings
	backticks := "`" + "`" + "`"
	return backticks + lang + "\n" + message + backticks
}

func isAdmin(discord *discordgo.Session, userID string, channelID string) bool {
	c, err := discord.Channel(channelID)
	if err != nil {
		return false
	}
	g, err := discord.Guild(c.GuildID)
	if err != nil {
		return false
	}
	//finish this later - look at DogBot on github for example
}

func commandChooser(discord *discordgo.Session, message *discordgo.MessageCreate) {
	command := strings.Fields(strings.ToLower(message.Content)) //note that channel names (hashtags) get converted to ID numbers, so this doesn't affect them
	command = append(command, "")
	fmt.Printf("Received command: %v", command)
	if strings.ToLower(command[0]) != (commandPrefix + strings.ToLower(botName)) {
		return
	}
	switch command[1] {
	case "":
		msg := `= Welcome To Polly! =

[ Commands ]
	- setup #channel1 #channel2 #channel3... :: takes a space-separated list of channels to learn from. If channels are unspecified, will use all of them.
	- dance :: a test command`
		discord.ChannelMessageSend(message.ChannelID, markdownWrapper("asciidoc", msg))
	case "setup":
		discord.ChannelMessageSend(message.ChannelID, "Setup!")
	case "dance":
		discord.ChannelMessageSend(message.ChannelID, ":dancer: Dancing! :dancer:") //w emojis
	default:
		discord.ChannelMessageSend(message.ChannelID, "Unrecognized command")
	}
	return
}
