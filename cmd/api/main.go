package main

import (
	"log"
	"net/http"
	"time"

	"github.com/krunalchavhan26/olx-api/internal/config"
	"github.com/krunalchavhan26/olx-api/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
