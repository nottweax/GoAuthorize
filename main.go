package main

import (
	"net/http"
	"rest/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
