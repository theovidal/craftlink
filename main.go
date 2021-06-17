package main

func main() {
	createBot()
	openWebsocket()
	defer ws.Close()

	go handleMinecraft()
	bot.Run(false)
}
