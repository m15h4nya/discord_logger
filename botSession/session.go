package botSession

import (
	hndlrs "discord_logger/botSession/handlers"
	"discord_logger/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Bot struct {
	Ready bool
	*discordgo.Session
	log *zap.SugaredLogger
}

func (b *Bot) CreateSession(cfg *config.Config, log *zap.SugaredLogger) {
	handler := &hndlrs.Handler{Cfg: cfg, OptState: ""}
	handlers := []interface{}{
		handler.MessageCreate,
		handler.MessageEdit,
		handler.MessageDelete,
		handler.MessageDeleteBulk,
		handler.Ready,
	}

	var err error
	b.Session, err = discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	b.Session.StateEnabled = false
	b.Ready = true
	hndlrs.AddHandlers(b.Session, handlers)
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
