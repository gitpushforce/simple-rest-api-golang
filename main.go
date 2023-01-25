package main

import (
	"fmt"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from go server")
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list of movies")
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie %s loaded", movie_id)
}
