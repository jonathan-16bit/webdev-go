package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	// Also catches other paths that dont have a more specific handler 
	mux.HandleFunc("/", home)  

	log.Println("Starting server on :4000")

	// Listen on port 4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
