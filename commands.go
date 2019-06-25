//Various commands

package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func commandChooser(discord *discordgo.Session, message *discordgo.MessageCreate) {
	command := strings.Fields(strings.ToLower(message.Content)) //note that channel names (hashtags) get converted to ID numbers, so this doesn't affect them
	command = append(command, "")
	fmt.Printf("Received command: %v", command)
	if strings.ToLower(command[0]) != (commandPrefix + strings.ToLower(botName)) {
		return
	}
	switch command[1] {
	case "":
		discord.ChannelMessageSend(message.ChannelID, "Help message")
	case "setup":
		discord.ChannelMessageSend(message.ChannelID, "Setup!")
	case "dance":
		discord.ChannelMessageSend(message.ChannelID, ":dancer: Dancing! :dancer:") //w emojis
	default:
		discord.ChannelMessageSend(message.ChannelID, "Unrecognized command")
	}
	return
}
