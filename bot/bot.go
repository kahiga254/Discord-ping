package bot

import (
	"discord-ping/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start (){
	var err error
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil{
		fmt.Println((err.Error()))
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println((err.Error()))
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return	
	}
	fmt.Println("Bot is running:")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){

	if m.Author.ID == BotID{
		return
	}

	if m.Content == "ping"{
		 s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
