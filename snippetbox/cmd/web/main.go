package main

import (
	"log"
	"net/http"
	"flag"
)

func main() {
	// Command-line arg 'addr', default value 4000 and some description
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()  // Actually reads into addr (without this, takes on default value)

	mux := http.NewServeMux()

	// Creates handler that sends the requested file back
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Removing the "/static" prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) 

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// flag.String() returns a pointer to the flag value, so we need to dereference it
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
