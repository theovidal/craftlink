package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/gorilla/websocket"
)

type Request struct {
	Command string `json:"command"`
	Token   string `json:"token"`
	Params  string `json:"params"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

var ws *websocket.Conn

var token string

func openWebsocket() {
	var err error
	ws, _, err = websocket.DefaultDialer.Dial(os.Getenv("WS_ADDRESS"), nil)
	if err != nil {
		log.Fatalln(Red.Sprint("‼ Couldn't connect to Minecraft server:", err))
	}
}

func handleMinecraft() {
	chatRegex := regexp.MustCompile("^<(.+)> (.+)$")

	for {
		var response Response
		if err := ws.ReadJSON(&response); err != nil && os.Getenv("ONYXCORD_ENV") == "development" {
			log.Println(Red.Sprint("‼ Couldn't read from Minecraft server:", err))
			return
		}
		switch response.Status {
		case 10:
			matches := chatRegex.FindSubmatch([]byte(response.Message))
			if matches == nil {
				break
			}

			username, content := matches[1], matches[2]

			if _, err := bot.Client.ChannelMessageSend(os.Getenv("CHANNEL_ID"), fmt.Sprintf("<%s> %s", username, content)); err != nil && os.Getenv("ONYXCORD_ENV") == "development" {
				log.Println(Red.Sprint("‼ Error sending the channel message:", err))
			}

			break
		case 200:
			token = response.Token
			break
		case 401:
			err := ws.WriteJSON(Request{
				Command: "LOGIN",
				Params:  os.Getenv("WS_PASSWORD"),
			})
			if err != nil {
				log.Fatalln(Red.Sprint("‼ Couldn't write to Minecraft server:", err))
			}
			break
		default:
			break
		}
	}
}
