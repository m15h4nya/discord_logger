package main

import (
	"discord_logger/configParser"
	dsHandlers "discord_logger/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	ready bool
	*discordgo.Session
}

func (b *Bot) CreateSession() {
	fmt.Println("FUCK")
	handler := &dsHandlers.Handler{Cfg: configParser.ParseConfig()}
	handlers := []interface{}{handler.MessageCreate, handler.MessageEdit, handler.MessageDelete}

	var err error
	b.Session, err = discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	b.StateEnabled = true
	b.ready = true
	b.State.MaxMessageCount = 500
	dsHandlers.AddHandlers(b.Session, handlers)
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		fmt.Println(err)
	}
}

func (b *Bot) StopSession() {
	b.ready = false
	err := b.Session.Close()
	if err != nil {
		fmt.Println(err)
	}
}
