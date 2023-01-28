package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

var collection = getSession().DB("go-study").C("movies")

var movies = Movies{
	Movie{"inception", 2010, "Nolan"},
	Movie{"fast and furious", 2015, "Diesel"},
	Movie{"Ninja turtles", 2014, "Liebesman"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from go server")
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "list of movies")
	json.NewEncoder(w).Encode(movies)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie %s loaded", movie_id)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movie_data)

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(movie_data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
