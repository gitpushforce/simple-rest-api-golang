package main

import (
	"log"
	"net/http"
)

func main() {
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/contacts", Contacts)
	// router.HandleFunc("/movies", MovieList)
	// router.HandleFunc("/movie/{id}", ShowMovie)
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
