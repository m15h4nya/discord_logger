package elastic

import (
	"context"
	"discord_logger/config"

	elasticlib "github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type Elastic struct {
	conn *elasticlib.Client
	cfg  *config.Config
	log  *zap.SugaredLogger
}

func NewElastic(cfg *config.Config, log *zap.SugaredLogger) (*Elastic, error) {
	conn, err := elasticlib.NewClient(
		elasticlib.SetSniff(false),
		elasticlib.SetURL(cfg.ElasticNodes...),
	)

	if err != nil {
		return nil, err
	}

	return &Elastic{conn: conn, cfg: cfg, log: log}, nil
}

func (e *Elastic) CreateUsersIndex(index string) error {

	_, err := e.conn.CreateIndex(index).BodyString(userIndex).Do(context.Background())
	if err != nil {
		elasticErr := err.(*elasticlib.Error)
		if elasticErr.Details.Type != ignoreError {
			return err
		}
	}

	return nil
}

func (e *Elastic) CreateMessagesIndex(index string) error {

	_, err := e.conn.CreateIndex(index).BodyString(messagesIndex).Do(context.Background())
	if err != nil {
		elasticErr := err.(*elasticlib.Error)
		if elasticErr.Details.Type != ignoreError {
			return err
		}
	}

	return nil
}
