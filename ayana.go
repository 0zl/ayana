package main

import (
	"log"

	"github.com/0zl/ayana/internal/config"
	"github.com/0zl/ayana/internal/server"
	"github.com/BurntSushi/toml"
)

func main() {
	cfg := config.DefaultConfig()

	_, err := toml.DecodeFile("config.toml", &cfg)
	if err != nil {
		log.Printf("Error decoding config file: %v", err)
	}

	srv := server.New()
	srv.SetupRoutes()

	log.Fatal(srv.Start(cfg.Server.Port))
}
