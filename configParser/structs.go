package configParser

type Config struct {
	Token             string   `json:"token"`
	LogChannelID      string   `json:"logChannelID"`
	GuildID           string   `json:"guildID"`
	IgnoreChannelsIDs []string `json:"channelsToIgnore"`
	Handlers          []string `json:"handlers"`
}
