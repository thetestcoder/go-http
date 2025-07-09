package main

import (
	"github.com/thetestcoder/basic-http/internal/api"
	"github.com/thetestcoder/basic-http/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load Config:", err)
	}

	api.NewServer(cfg)
}
