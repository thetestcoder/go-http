package api

import (
	"fmt"
	"github.com/thetestcoder/basic-http/internal/config"
	"log"
	"net/http"
	"time"
)

func NewServer(cfg *config.Config) {
	h := &http.Server{
		Handler:      setupRouter(),
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	log.Println("Server start on port " + cfg.Server.Port)
	log.Fatal(h.ListenAndServe())

}
