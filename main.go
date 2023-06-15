package main

import (
	"github.com/arohan007/gameModeStats/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	//r.HandleFunc("/stats", handleStats).Methods(http.MethodPost)
	//r.HandleFunc("/popularity", handlePopularity).Methods(http.MethodGet)
	//r.HandleFunc("/createUser", createUser).Methods(http.MethodPost)
	r.HandleFunc("/hello", Controller.DummyHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
