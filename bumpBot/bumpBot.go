package bumpBot

import (
	"context"
	"discord_logger/configParser"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Bump() {
	cfg := configParser.ParseConfig()

	b, err := discordgo.New(cfg.BumpToken)
	if err != nil {
		fmt.Println(err)
	}
	b.Open()
	if err != nil {
		fmt.Println(err)
	}

	for {
		ctx, _ := context.WithTimeout(context.TODO(), 5*time.Minute)
		<-ctx.Done()
		_, err = b.ChannelMessageSend(cfg.BumpChannelID, "!bump")
		if err != nil {
			fmt.Println(err)
		}
	}
}
