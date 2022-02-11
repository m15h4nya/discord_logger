package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type token struct {
	Token string
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "It's working")
	}
}

func main() {
	token := token{}
	file, err := os.Open("./config/config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	js, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &token)

	if err != nil {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + token.Token)
	discord.AddHandler(messageCreate)

	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	if err != nil {
		log.Fatal(err)
	}
}
