package http

import (
	"discord_logger/botSession"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Service struct {
	//serverErrorChannel chan error
	router     *mux.Router
	httpServer *http.Server
	//sig        chan bool
	botSession *botSession.Bot
}

func (h *Service) InitService() {
	//h.sig = make(chan bool, 1)
	h.botSession = &botSession.Bot{}
	//h.router = mux.NewRouter()
	h.httpServer = &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	h.botSession.CreateSession()
	h.botSession.StartSession()
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
		h.botSession.CreateSession()
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
		/*case "/audit":
		if !h.botSession.Ready {
			_, err := fmt.Fprintf(w, "the bot is stopped")
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		audit, _ := h.botSession.GuildAuditLog("465780328611708937", "", "", 72, 100)
		res := ""
		for _, v := range audit.AuditLogEntries {
			t, _ := discordgo.SnowflakeTimestamp(v.ID)
			res += t.String() + "\n"
		}
		_, _ = fmt.Fprintf(w, "%v", res)*/
	}
}

func (h *Service) CreateServer() {
	err := h.httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
