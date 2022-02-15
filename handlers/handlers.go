package handlers

import (
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	logChannel = configParser.ParseLogChannel()
	guild      = configParser.ParseGuild()
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {

}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "test" {
		_, err := s.ChannelMessageSend(logChannel.ID, "It's working")
		if err != nil {
			log.Printf("MessageCreate: %v", err)
		}
	}
}

func MessageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.BeforeUpdate == nil || m.Content == "" || m.BeforeUpdate.Author.ID == s.State.User.ID {
		return
	}
	msgAuthor := m.BeforeUpdate.Author.Username
	msgOldContent := m.BeforeUpdate.Content
	msgNewContent := m.Content
	_, err := s.ChannelMessageSend(logChannel.ID, msgAuthor+": "+msgOldContent+" -> "+msgNewContent)
	if err != nil {
		log.Printf("MessageEdit: %v\n", err)
	}
}

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	if m.BeforeDelete == nil || m.BeforeDelete.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.BeforeDelete.ID)
	auditLog, err := s.GuildAuditLog(guild.ID, "", "", int(discordgo.AuditLogActionMessageDelete), 100)
	if err != nil {
		fmt.Printf("MessageDelete on \"auditLog, err :=...\" : %v\n", err)
	}
	msgAuthor := m.BeforeDelete.Author.Username
	msgContent := m.BeforeDelete.Content
	eventAuthor := m.BeforeDelete.Author.Username
	for _, entry := range auditLog.AuditLogEntries {
		if entry.TargetID == m.ID {
			user, _ := s.User(entry.UserID)
			eventAuthor = user.Username
		}
	}

	_, err = s.ChannelMessageSend(logChannel.ID, eventAuthor+": **deleted "+msgAuthor+"'s message** -> "+msgContent)
	if err != nil {
		log.Printf("MessageDelete on \"_, err = s.ChannelMessageSend(...\" : %v\n", err)
	}
}
