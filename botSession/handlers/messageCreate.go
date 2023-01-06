package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const chars = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯабвгдеёжзийклмнопрстуфхцчшщьыъэюя"
const mainChat = "1030029993155239956"
const mainChatRu = "796301641862086676"

func (h *Handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ContainsAny(m.Message.Content, chars) && m.ChannelID == mainChat {
		if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
			log.Printf("MessageCreate: %v\n", err)
		}
		msg := fmt.Sprintf("<@%s>, cyrillic is baned here, go to <#%s>", m.Author.ID, mainChatRu)
		s.ChannelMessageSend(mainChat, msg)
	}

	err := s.State.MessageAdd(m.Message)
	if err != nil {
		log.Printf("MessageCreate: %v\n", err)
	}

}
