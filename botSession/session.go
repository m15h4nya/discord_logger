package botSession

import (
	handlers2 "discord_logger/botSession/handlers"
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	Ready bool
	*discordgo.Session
}

func (b *Bot) CreateSession() {
	handler := &handlers2.Handler{Cfg: configParser.ParseConfig(), OptState: ""}
	handlers := []interface{}{handler.MessageCreate, handler.MessageEdit, handler.MessageDelete}

	var err error
	b.Session, err = discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	b.StateEnabled = true
	b.Ready = true
	b.State.MaxMessageCount = 500
	handlers2.AddHandlers(b.Session, handlers)
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
