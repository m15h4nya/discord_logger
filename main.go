package main

import (
	"discord_logger/config"
	"discord_logger/logger"
	"discord_logger/service"
)

func main() {
	server := service.Service{}
	log := logger.NewLogger()
	cfg := config.ParseConfig()
	server.InitService(cfg, log)
	server.StartHTTP()
}
