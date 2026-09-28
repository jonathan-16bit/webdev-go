package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Creates handler that sends the requested file back
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Removing the "/static" prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) 

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetView)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
