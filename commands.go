//Various commands

package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func markdownWrapper(lang string, message string) string {
	//hack for working with backticks in go strings
	backticks := "`" + "`" + "`"
	return backticks + lang + "\n" + message + backticks
}

func isAdmin(discord *discordgo.Session, userID string, channelID string) bool {
	//return false if err, err on the side of caution.
	c, err := discord.Channel(channelID)
	if err != nil {
		return false
	}
	g, err := discord.Guild(c.GuildID)
	if err != nil {
		return false
	}
	//Although owners may not *technically* be admins, they do have similar priviledges
	//Not technically correct, but I think this should stay
	if g.OwnerID == userID {
		return true
	}

	member, err := discord.GuildMember(g.ID, userID)
	if err != nil {
		return false
	}

	roles := member.Roles
	for i := 0; i < len(roles); i++ {
		role, _ := discord.State.Role(g.ID, roles[i])
		if (role.Permissions & discordgo.PermissionAdministrator) == discordgo.PermissionAdministrator {
			return true
		}
	}
	return false
}

func setup(discord *discordgo.Session, command []string) ([]string, error) {
	//get message history of each channel
	messagesPerChannel := 100 //max # of messages per channel. arbitrary, to be turned into a env variable later
	channels := command[2:]

	chatlogs := []string{}

	for i := 0; i < len(channels)-1; i++ { //the -1 is cause we append an empty string earlier
		channelID := cleanChannelId(channels[i])
		fmt.Println(channelID)
		messages, err := discord.ChannelMessages(channelID, messagesPerChannel, "", "", "")
		if err != nil {
			continue
		}
		for j := 0; j < len(messages); j++ {
			content := messages[j].Content
			first := string(content[0])
			author := messages[j].Author.ID
			isProbablyBotCommand, _ := regexp.MatchString("[!$%^&*,.?:{}|<>`]", first)
			if messages[j].Author.ID != botID && !isProbablyBotCommand {
				fmt.Printf("%s: %s, \n", author, content)
				chatlogs = append(chatlogs, content)
			}
		}
	}
	return chatlogs, nil
}

func cleanChannelId(channelID string) string {
	//converts <#xyz> id to 'xyz'
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return reg.ReplaceAllString(channelID, "")
}

func commandChooser(discord *discordgo.Session, message *discordgo.MessageCreate) {
	command := strings.Fields(strings.ToLower(message.Content)) //note that channel names (hashtags) get converted to ID numbers, so this doesn't affect them
	command = append(command, "")
	fmt.Printf("Received command: %v \n", command)
	if strings.ToLower(command[0]) != (commandPrefix + strings.ToLower(botName)) {
		return
	}
	switch command[1] {
	case "":
		msg := `= Welcome To Polly! =

[ Commands ]
	Admin Commands:
		- setup #channel1 #channel2 #channel3... :: takes a space-separated list of channels to learn from. If channels are unspecified, will use all of them.
		- setmode silent/normal/chatty :: silent prevents the bot from speaking unless specifically invoked. normal/chatty allow the bot to speak randomly in the chat, but degree varies bassed on mode.
	User Commands:
		- dance :: a test command`
		discord.ChannelMessageSend(message.ChannelID, markdownWrapper("asciidoc", msg))
	case "setup":
		admin := isAdmin(discord, message.Author.ID, message.ChannelID)
		if admin {

			func() {
				discord.ChannelMessageSend(message.ChannelID, "Sure thing, you're an admin!")
				go setup(discord, command)
				defer discord.ChannelMessageSend(message.ChannelID, ":bird: All set up :bird:")
			}()

		} else {
			discord.ChannelMessageSend(message.ChannelID, "You must have an administrator role to use this command.")
		}
	case "setmode":
		admin := isAdmin(discord, message.Author.ID, message.ChannelID)
		if admin {
			discord.ChannelMessageSend(message.ChannelID, "Sure thing, you're an admin!")
		} else {
			discord.ChannelMessageSend(message.ChannelID, "You must have an administrator role to use this command.")
		}
	case "dance":
		discord.ChannelMessageSend(message.ChannelID, ":dancer: Dancing! :dancer:") //w emojis
	default:
		discord.ChannelMessageSend(message.ChannelID, "Unrecognized command")
	}
	return
}
