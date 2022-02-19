package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "test" {
		_, err := s.ChannelMessageSend(h.Cfg.LogChannelID, "It's working")
		if err != nil {
			log.Printf("MessageCreate: %v", err)
		}
	}
}
