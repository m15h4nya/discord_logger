package handlers

import (
	"discord_logger/errorLogger"
	"github.com/bwmarrin/discordgo"
)

var (
	logChannel = discordgo.Channel{ID: "531632649526050822"}
	Messages   = make(map[string]*discordgo.Message)
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	/*
		guild, _ := s.Guild("465780328611708937")
		s.State.GuildAdd(guild)
		ch, _ := s.State.Channel("531632649526050822")
		fmt.Printf("%#v", ch.Messages)
	*/
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Messages[m.ID] = m.Message

	if m.Content == "test" {
		_, err := s.ChannelMessageSend(logChannel.ID, "It's working")
		errorLogger.CheckErr(err, "handlers.MessageCreate")
	}
}

func MessageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	defer func() {
		_ = recover()
	}()
	msgAuthor := m.BeforeUpdate.Author.Username
	msgOldContent := m.BeforeUpdate.Content
	msgNewContent := m.Content
	// Messages[m.ID] = nil
	_, err := s.ChannelMessageSend(logChannel.ID, msgAuthor+": "+msgOldContent+" -> "+msgNewContent)
	if err != nil {
		errorLogger.CheckErr(err, "handlers.MessageDelete")
	}
}

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	defer func() {
		_ = recover()
	}()
	/*
		msgAuthor := Messages[m.ID].Author.Username
		msgContent := Messages[m.ID].Content
		Messages[m.ID] = nil
	*/
	msgAuthor := m.BeforeDelete.Author.Username
	msgContent := m.BeforeDelete.Content
	_, err := s.ChannelMessageSend(logChannel.ID, msgAuthor+": **deleted message** -> "+msgContent)
	if err != nil {
		errorLogger.CheckErr(err, "handlers.MessageDelete")
	}
}
