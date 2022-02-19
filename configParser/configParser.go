package configParser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ParseConfig() (cfg Config) {
	cfg = Config{}
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	js, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &cfg)

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
