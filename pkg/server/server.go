package server

import (
	"bookcrossing-backend/pkg/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Launch(r *mux.Router, cfg *config.Config) {
	fmt.Printf("Starting server at %s\n", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatal(err)
	}
}
