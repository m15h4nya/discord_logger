package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
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
