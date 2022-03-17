package main

import "discord_logger/bumpBot"

func main() {
	go bumpBot.Bump()
	server := HTTPService{}
	server.InitService()
	server.createServer()
}
