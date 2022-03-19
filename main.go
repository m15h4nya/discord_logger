package main

import (
	"discord_logger/bumpBot"
	"discord_logger/http"
)

func main() {
	go bumpBot.Bump()
	server := http.Service{}
	server.InitService()
	server.CreateServer()
}
