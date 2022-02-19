package main

import (
	"discord_logger/configParser"
	dsHandlers "discord_logger/handlers"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
