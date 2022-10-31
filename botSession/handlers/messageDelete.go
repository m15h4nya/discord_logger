package handlers

import (
	"discord_logger/config"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	msg, err := s.State.Message(m.ChannelID, m.ID)
	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}

	if msg == nil || msg.Author.ID == s.State.User.ID ||
		config.Contains(m.ChannelID, h.Cfg.IgnoreChannelsIDs) {
		return
	}

	msgAuthor := msg.Author.Username
	msgContent := msg.Content
	msgAttachments := msg.Attachments
	msgChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}

	eventAuthor := msg.Author.Username
	auditLog, err := s.GuildAuditLog(m.GuildID, "", "", int(discordgo.AuditLogActionMessageDelete), 1)
	if auditLog.AuditLogEntries[0].ID != h.OptState {
		t, _ := s.User(auditLog.AuditLogEntries[0].UserID)
		eventAuthor = t.Username
		h.OptState = auditLog.AuditLogEntries[0].ID
	}

	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}

	logAttachmentsMsg := ""
	if len(msgAttachments) != 0 {
		logAttachmentsMsg += "Attachments:\n"
		for _, attachment := range msgAttachments {
			aURL := fmt.Sprintf("%v\n", attachment.URL)
			logAttachmentsMsg += aURL
		}
	}

	logMsg := fmt.Sprintf("`%v: deleted %v's message in %v` -> %v\n", eventAuthor, msgAuthor, msgChannel.Name, msgContent)

	_, err = s.ChannelMessageSend(h.Cfg.LogChannelID, logMsg+logAttachmentsMsg)
	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}

	err = s.State.MessageRemove(m.Message)
	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}
}
