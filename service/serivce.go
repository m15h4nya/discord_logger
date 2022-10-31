package service

import (
	"discord_logger/botSession"
	"discord_logger/config"
	"discord_logger/elastic"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	cfg        *config.Config
	log        *zap.SugaredLogger
	httpServer *http.Server
	botSession *botSession.Bot
	elastic    *elastic.Elastic
}

func (h *Service) InitService(cfg *config.Config, log *zap.SugaredLogger) (err error) {

	h.cfg = cfg
	h.botSession = &botSession.Bot{}

	h.elastic, err = elastic.NewElastic(cfg, log)
	if err != nil {
		log.Error(err)
		return err
	}

	h.httpServer = &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	h.elastic.CreateMessagesIndex(cfg.Elastic.MessagesIndex)
	h.elastic.CreateUsersIndex(cfg.Elastic.UsersIndex)

	h.botSession.CreateSession(cfg, log)
	h.botSession.StartSession()

	return nil
}

func (h Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		_, err := fmt.Fprintf(w, "basePage")
		if err != nil {
			fmt.Println(err)
		}

	case "/start":
		if h.botSession.Ready {
			_, err := fmt.Fprintf(w, "already started")
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		h.botSession.CreateSession(h.cfg, h.log) // creating new session so it can start again
		h.botSession.StartSession()

		_, err := fmt.Fprintf(w, "startPage")
		if err != nil {
			fmt.Println(err)
		}

	case "/stop":
		if !h.botSession.Ready {
			_, err := fmt.Fprintf(w, "already stopped")
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		h.botSession.StopSession()

		_, err := fmt.Fprintf(w, "stopPage")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Service) StartHTTP() {
	err := h.httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
