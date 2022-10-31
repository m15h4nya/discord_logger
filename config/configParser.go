package config

import (
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Logger
	Elastic
}

type Logger struct {
	Token             string   `toml:"token"`
	LogChannelID      string   `toml:"log_id"`
	GuildID           string   `toml:"guild_id"`
	IgnoreChannelsIDs []string `toml:"ignore_channels"`
}

type Elastic struct {
	ElasticNodes    []string `toml:"nodes"`
	MessagesIndex   string   `toml:"messages_index"`
	UsersIndex      string   `toml:"users_index"`
	MaxMessageCount int      `toml:"msg_count"`
}

func ParseConfig() (cfg *Config) {
	cfg = &Config{}
	file, err := os.ReadFile("./config/config.toml")

	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(file, cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func Contains(str string, sl []string) bool {
	for _, v := range sl {
		if str == v {
			return true
		}
	}
	return false
}
