package botSession

import (
	"discord_logger/botSession/handlers"
	"discord_logger/configParser"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Ready bool
	*discordgo.Session
}

func (b *Bot) CreateSession() {
	handler := handlers.NewHandler(configParser.ParseConfig())
	hndlrs := []interface{}{handler.MessageCreate, handler.MessageEdit, handler.MessageDelete, handler.MessageDeleteBulk, handler.Ready}

	var err error
	b.Session, err = discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	b.Session.StateEnabled = false
	b.Ready = true
	handlers.AddHandlers(b.Session, hndlrs)
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		fmt.Println(err)
	}
}

func (b *Bot) StopSession() {
	b.Ready = false
	err := b.Session.Close()
	if err != nil {
		fmt.Println(err)
	}
}
