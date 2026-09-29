package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// Command-line arg 'addr', default value 4000 and some description
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()  // Actually reads into addr (without this, takes on default value)

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Lshortfile: file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	// Creates handler that sends the requested file back
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Removing the "/static" prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) 

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	// flag.String() returns a pointer to the flag value, so we need to dereference it
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
