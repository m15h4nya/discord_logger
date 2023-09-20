package handlers

import (
	"fmt"
	"log"
	"strings"
	"time"

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

		if _, ok := h.warnedUsers[m.Author.ID]; !ok {
			msg := fmt.Sprintf("<@%s>, this chat is english only, go to <#%s>", m.Author.ID, mainChatRu)
			s.ChannelMessageSend(mainChat, msg)
		}

		h.warnedUsers[m.Author.ID] += 1
		if h.warnedUsers[m.Author.ID] == 5 {
			until := time.Now().Add(time.Minute * 1)
			s.GuildMemberTimeout(m.GuildID, m.Author.ID, &until)
			delete(h.warnedUsers, m.Author.ID)
		}

	}

	err := s.State.MessageAdd(m.Message)
	if err != nil {
		log.Printf("MessageCreate: %v\n", err)
	}

}
