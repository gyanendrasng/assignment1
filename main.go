package main

import (
	"net/http"
	"time"
)

type Config struct {
	ListenAddress string
}

type Server struct {
	Config
}

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Addr:         ":3001",
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	srv.ListenAndServe()
}
