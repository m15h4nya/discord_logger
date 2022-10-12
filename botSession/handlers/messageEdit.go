package handlers

import (
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	msg, err := s.State.Message(m.ChannelID, m.ID)
	if err != nil {
		log.Printf("MessageDelete: %v\n", err)
	}

	if msg == nil || m.Content == "" || msg.Author.ID == s.State.User.ID ||
		configParser.Contains(m.ChannelID, h.Cfg.IgnoreChannelsIDs) {
		return
	}

	msgAuthor := msg.Author.Username
	msgChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Printf("MessageEdit: %v\n", err)
	}

	msgOldContent := msg.Content
	msgOldAttachments := msg.Attachments
	msgNewContent := m.Content
	msgNewAttachments := m.Attachments

	logAttachmentsMsg := ""
	if len(msgOldAttachments) != 0 || len(msgNewAttachments) != 0 {
		logAttachmentsMsg += "Old attachments:\n"
		for _, attachment := range msgOldAttachments {
			aURL := fmt.Sprintf("%v\n", attachment.URL)
			logAttachmentsMsg += aURL
		}
		logAttachmentsMsg += "New attachments:\n"
		for _, attachment := range msgNewAttachments {
			aURL := fmt.Sprintf("%v\n", attachment.URL)
			logAttachmentsMsg += aURL
		}
	}

	logMsg := fmt.Sprintf("`%v: edited message in %v` %v -> %v\n", msgAuthor, msgChannel.Name, msgOldContent, msgNewContent)
	_, err = s.ChannelMessageSend(h.Cfg.LogChannelID, logMsg+logAttachmentsMsg)
	if err != nil {
		log.Printf("MessageEdit: %v\n", err)
	}

	err = s.State.MessageAdd(m.Message)
	if err != nil {
		log.Printf("MessageEdit: %v\n", err)
	}

}
