package main

import (
	"log"

	"example.com/m/internal/config"
	"example.com/m/internal/database"
	"example.com/m/internal/server"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	db, err := database.New(&cfg)
	if err != nil {
		log.Fatalf("database error: %s", err)
	}
	defer db.Close()

	s, err := server.New(&cfg)
	if err != nil {
		log.Fatalf("server error: %s", err)
	}
	defer s.Close()

	s.Listen()
}
