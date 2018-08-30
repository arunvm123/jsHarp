package operations

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	m := mux.NewRouter()

	m.Handle("/code", retrieveCode()).Methods("POST")
	m.Handle("/render", renderPage()).Methods("GET")

	return m
}
