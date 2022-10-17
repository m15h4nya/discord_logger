package handlers

import (
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageDeleteBulk(s *discordgo.Session, m *discordgo.MessageDeleteBulk) {

	for i := 0; i < len(m.Messages)/2; i++ {
		j := len(m.Messages) - i - 1
		m.Messages[i], m.Messages[j] = m.Messages[j], m.Messages[i]
	}

	for _, message := range m.Messages {
		msg, err := s.State.Message(m.ChannelID, message)
		if err != nil {
			log.Printf("MessageDeleteBulk: %v\n", err)
			continue
		}
		if msg == nil || msg.Author.ID == s.State.User.ID ||
			configParser.Contains(m.ChannelID, h.Cfg.IgnoreChannelsIDs) {
			continue
		}

		msgAuthor := msg.Author.Username
		msgContent := msg.Content
		msgAttachments := msg.Attachments
		msgChannel, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Printf("MessageDeleteBulk: %v\n", err)
		}

		eventAuthor := "BULK DELETE"
		auditLog, err := s.GuildAuditLog(m.GuildID, "", "", int(discordgo.AuditLogActionMessageBulkDelete), 1)
		if auditLog.AuditLogEntries[0].ID != h.OptStateBulk {
			t, _ := s.User(auditLog.AuditLogEntries[0].UserID)
			eventAuthor = t.Username
			h.OptStateBulk = auditLog.AuditLogEntries[0].ID
		}

		if err != nil {
			log.Printf("MessageDeleteBulk: %v\n", err)
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
			log.Printf("MessageDeleteBulk: %v\n", err)
		}

		//messages = append(messages, *msg)

		err = s.State.MessageRemove(msg)
		if err != nil {
			log.Printf("MessageDeleteBulk: %v\n", err)
		}
	}

}
