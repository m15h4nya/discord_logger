package handlers

import (
	"discord_logger/configParser"
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Cfg          configParser.Config
	OptState     string
	OptStateBulk string
}

func (h *Handler) removePings(s *discordgo.Session, m *discordgo.Message) (content string) {
	content = m.Content

	userRegex := regexp.MustCompile(`<@\d*>`)
	channelsRegex := regexp.MustCompile(`<#\d*>`)
	rolesRegex := regexp.MustCompile(`<@&\d*>`)

	usersIds := userRegex.FindAll([]byte(m.Content), -1)
	channelsIds := channelsRegex.FindAll([]byte(m.Content), -1)
	rolesIds := rolesRegex.FindAll([]byte(m.Content), -1)

	for _, v := range usersIds {
		user, err := s.User(string(v)[2 : len(v)-1])
		if err != nil {
			continue
		}
		content = strings.ReplaceAll(content,
			string(v),
			user.Username+"#"+user.Discriminator,
		)
	}

	for _, v := range channelsIds {
		channel, err := s.Channel(string(v)[2 : len(v)-1])
		if err != nil {
			continue
		}
		content = strings.ReplaceAll(content,
			string(v),
			"#"+channel.Name,
		)
	}

	for _, v := range rolesIds {
		role, err := s.State.Role(m.GuildID, (string(v)[3 : len(v)-1]))
		if err != nil {
			continue
		}
		content = strings.ReplaceAll(content,
			string(v),
			"@"+role.Name,
		)
	}

	return content
}
