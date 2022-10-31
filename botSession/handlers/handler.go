package handlers

import (
	"discord_logger/config"
)

type Handler struct {
	Cfg          *config.Config
	OptState     string
	OptStateBulk string
}
