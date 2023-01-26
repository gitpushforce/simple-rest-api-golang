package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacts", Contacts)
	router.HandleFunc("/movies", MovieList)
	router.HandleFunc("/movie/{id}", ShowMovie)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
