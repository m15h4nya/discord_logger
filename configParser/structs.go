package configParser

type Token struct {
	Token string `json:"token"`
}

type LogChannel struct {
	ID string `json:"logChannelID"`
}

type Guild struct {
	ID string `json:"guildID"`
}
