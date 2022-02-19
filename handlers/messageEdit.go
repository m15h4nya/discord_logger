package handlers

import (
	"discord_logger/configParser"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.BeforeUpdate == nil || m.Content == "" || m.BeforeUpdate.Author.ID == s.State.User.ID ||
		configParser.Contains(m.ChannelID, h.Cfg.IgnoreChannelsIDs) {
		return
	}
	msgAuthor := m.BeforeUpdate.Author.Username
	msgOldContent := m.BeforeUpdate.Content
	msgNewContent := m.Content
	_, err := s.ChannelMessageSend(h.Cfg.LogChannelID, msgAuthor+": "+msgOldContent+" -> "+msgNewContent)
	if err != nil {
		log.Printf("MessageEdit: %v\n", err)
	}
}
