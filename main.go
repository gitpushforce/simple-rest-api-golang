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

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from go server")
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}
