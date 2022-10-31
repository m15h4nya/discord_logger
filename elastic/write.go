package elastic

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type Message struct {
	Id          string     `json:"id"`
	ChannelId   string     `json:"channel_id"`
	GuildId     string     `json:"guild_id"`
	Content     string     `json:"content"`
	Attachment  Attachment `json:"attachment"`
	Author      string     `json:"author"`
	ReferenceId string     `json:"reference_id"`
}

type Attachment struct {
	FileName      string `json:"file_name"`
	Bytes         string `json:"bytes"`
	FileExtension string `json:"file_extension"`
}

func (e *Elastic) WriteMessage(msg Message) (err error) {

	// body, err := json.Marshal(msg)
	// if err != nil {
	// 	return err
	// }

	if _, err = e.conn.Index().Index(e.cfg.Elastic.MessagesIndex).BodyJson(msg).Do(context.Background()); err != nil {
		e.log.Errorf("WriteMessage: %s", err)
		return err
	}
	return nil
}
