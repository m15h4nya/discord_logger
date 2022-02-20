package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPService struct {
	//serverErrorChannel chan error
	router     *mux.Router
	httpserver *http.Server
}

func (h *HTTPService) InitService() {
	h.router = mux.NewRouter()
	h.router.Handle("/", h)
	h.httpserver = &http.Server{
		Addr:              ":8080",
		Handler:           h.router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
}

func (h *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "basePage")
	if err != nil {
		fmt.Println(err)
	}
}

func createServer() {
	service := HTTPService{}
	service.InitService()
	service.httpserver.ListenAndServe()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
