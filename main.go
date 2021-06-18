package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var members map[string]string

func main() {
	data, err := ioutil.ReadFile("members.yml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &members)
	if err != nil {
		log.Fatal(Red.Sprint("‼ Error reading the members file: ", err))
	}

	createBot()
	openWebsocket()
	defer ws.Close()

	go handleMinecraft()
	bot.Run(false)
}
