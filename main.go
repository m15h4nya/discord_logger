package main

import (
	"discord_logger/http"
)

func main() {
	server := http.Service{}
	server.InitService()
	server.CreateServer()
}
