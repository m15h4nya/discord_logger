package main

import (
	"discord_logger/configParser"
	dsHandlers "discord_logger/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func createSession(c chan bool) {
	handler := &dsHandlers.Handler{Cfg: configParser.ParseConfig()}
	handlers := []interface{}{handler.MessageCreate, handler.MessageEdit, handler.MessageDelete}

	discord, err := discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	discord.StateEnabled = true
	discord.State.MaxMessageCount = 500
	dsHandlers.AddHandlers(discord, handlers)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	if <-c {
		err := discord.Close()
		if err != nil {
			fmt.Printf("Error while closing the session: %v", err)
		}
		return
	}
}
