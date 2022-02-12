package configParser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type token struct {
	Token string
}

func ParseToken() (Token string) {
	token := token{}
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

	return token.Token
}
