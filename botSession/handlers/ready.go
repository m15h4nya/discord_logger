package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) Ready(s *discordgo.Session, m *discordgo.Ready) {
	s.State = discordgo.NewState()
	s.State.Ready = *m
	s.State.MaxMessageCount = 500

	guild, err := s.Guild(h.Cfg.GuildID)
	channels, err := s.GuildChannels(h.Cfg.GuildID)
	if err != nil {
		log.Printf("Ready: %v\n", err)
		return
	}
	err = s.State.GuildAdd(guild)
	if err != nil {
		log.Printf("Ready: %v\n", err)
		return
	}
	for _, channel := range channels {
		err = s.State.ChannelAdd(channel)
		if err != nil {
			log.Printf("Ready: %v\n", err)
		}
	}
}
