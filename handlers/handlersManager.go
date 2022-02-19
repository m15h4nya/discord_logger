package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func AddHandlers(s *discordgo.Session, handlers []interface{}) {
	for _, handler := range handlers {
		s.AddHandler(handler)
	}
}
