package handlers

import (
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	if m.BeforeDelete == nil || m.BeforeDelete.Author.ID == s.State.User.ID ||
		configParser.Contains(m.ChannelID, h.Cfg.IgnoreChannelsIDs) {
		return
	}
	//auditLog, err := s.GuildAuditLog(h.Cfg.GuildID, "", "", int(discordgo.AuditLogActionMessageDelete), 100)
	//if err != nil {
	//	fmt.Printf("MessageDelete on \"auditLog, err :=...\" : %v\n", err)
	//}
	msgAuthor := m.BeforeDelete.Author.Username
	msgContent := m.BeforeDelete.Content
	eventAuthor := m.BeforeDelete.Author.Username
	msgAttachments := m.Attachments
	msg := fmt.Sprintf("%v: **deleted %v's message** -> %v \nAttachments -> %v", eventAuthor, msgAuthor, msgContent, msgAttachments)
	//for _, entry := range auditLog.AuditLogEntries {
	//	if entry.TargetID == m.ID {
	//		user, _ := s.User(entry.UserID)
	//		eventAuthor = user.Username
	//	}
	//	fmt.Println(discordgo.SnowflakeTimestamp(entry.TargetID))
	//}

	_, err := s.ChannelMessageSend(h.Cfg.LogChannelID, msg)
	if err != nil {
		log.Printf("MessageDelete on \"_, err = s.ChannelMessageSend(...\" : %v\n", err)
	}
}
