package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	err := s.State.MessageAdd(m.Message)
	if err != nil {
		log.Printf("MessageCreate: %v\n", err)
	}

}
