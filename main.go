package main

import(
	"fmt"
	"discord-ping/config"
	"discord-ping/bot"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}