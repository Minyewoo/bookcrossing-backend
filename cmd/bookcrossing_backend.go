package main

import (
	"bookcrossing-backend/pkg/config"
	"bookcrossing-backend/pkg/server"
	"bookcrossing-backend/pkg/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.ReadConfig()

	dbClient := storage.SetupStorage(&cfg.DBCredentials)
	defer dbClient.Close()

	router := server.SetupRouter(dbClient)

	server.Launch(router, cfg)
}
