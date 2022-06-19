package server

import (
	"bookcrossing-backend/ent"

	"github.com/gorilla/mux"
)

func SetupRouter(client *ent.Client) *mux.Router {
	r := mux.NewRouter()
	return r
}
