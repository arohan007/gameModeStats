package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

userData := Dal.NewInitializeUser()

func main() {



	r := mux.NewRouter()
	r.HandleFunc("/stats", handleStats).Methods(http.MethodPost)
	r.HandleFunc("/popularity", handlePopularity).Methods(http.MethodGet)
	r.HandleFunc("/createUser", createUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
