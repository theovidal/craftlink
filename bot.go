package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var bot onyxcord.Bot

func createBot() {
	bot = onyxcord.RegisterBot("craftlink")
	bot.Client.AddHandler(handleMessage)
}

func handleMessage(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot || message.ChannelID != os.Getenv("CHANNEL_ID") {
		return
	}

	username := message.Author.Username
	if message.Member.Nick != "" {
		username = message.Member.Nick
	}

	output := strings.ReplaceAll(message.Content, "\n", " ")
	for _, file := range message.Attachments {
		output += "\n" + file.ProxyURL
	}
	if output == "" {
		return
	}

	outputParts, partIndex := []string{""}, 0
	for _, word := range strings.Split(output, " ") {
		messageWithWord := outputParts[partIndex] + " " + word
		if (partIndex == 0 && len(messageWithWord) > 236) || (partIndex > 0 && len(messageWithWord) > 256) {
			partIndex++
			outputParts = append(outputParts, "")
		}
		outputParts[partIndex] += " " + word
	}

	for index, part := range outputParts {
		content := fmt.Sprintf(`{"text":"%s", "color": "white"}`, part)
		if index == 0 {
			content = fmt.Sprintf(`{"text":"[Discord]","color":"blue"},{"text":" <%s>", "color": "white"},%s`, username, content)
		}

		err := ws.WriteJSON(Request{
			Command: "EXEC",
			Token:   token,
			Params:  fmt.Sprintf(`tellraw @a [%s]`, content),
		})
		if err != nil {
			log.Println("write:", err)
		}
	}
}
