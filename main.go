package main

import (
	"discord_logger/configParser"
	"discord_logger/handlers"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := configParser.ParseToken()
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	discord.StateEnabled = true
	discord.State.MaxMessageCount = 100
	discord.AddHandler(handlers.Ready)
	discord.AddHandler(handlers.MessageCreate)
	discord.AddHandler(handlers.MessageDelete)
	discord.AddHandler(handlers.MessageEdit)

	err = discord.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	if err != nil {
		log.Fatal(err)
	}
}
