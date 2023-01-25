package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world from go server")
	})

	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)
}
