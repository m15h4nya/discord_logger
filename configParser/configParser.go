package configParser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ParseToken() (token Token) {
	token = Token{}
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	js, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &token)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

func ParseLogChannel() (logChannel LogChannel) {
	logChannel = LogChannel{}
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	js, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &logChannel)

	if err != nil {
		log.Fatal(err)
	}

	return logChannel
}

func ParseGuild() (guild Guild) {
	guild = Guild{}
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	js, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &guild)

	if err != nil {
		log.Fatal(err)
	}

	return guild
}
