package handlers

import (
	"discord_logger/configParser"
)

type Handler struct {
	Cfg      configParser.Config
	OptState string
}
