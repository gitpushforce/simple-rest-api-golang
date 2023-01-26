package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from go server")
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"inception", 2010, "Nolan"},
		Movie{"fast and furious", 2015, "Diesel"},
		Movie{"Ninja turtles", 2014, "Liebesman"},
	}

	//fmt.Fprintf(w, "list of movies")
	json.NewEncoder(w).Encode(movies)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie %s loaded", movie_id)
}
