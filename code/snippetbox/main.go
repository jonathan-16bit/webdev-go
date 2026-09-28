package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Restricting the root url pattern
	// Returns "404 page not found" when URL path isn't exactly "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hi from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	/*
	 * We can only call w.WriteHeader() once per response, and after the status 
	 * code is written, it cannot be changed.
	 *
	 * If we don't call w.WriteHeader() explicitly, then the first call to w.Write() 
	 * will automatically send a 200 OK status code.
	 * So, if we want to send a non-200 status code, we must call w.WriteHeader() before 
	 * ANY call to w.Write()
	*/
	if r.Method != "POST" {
		// Add a new header "Allow" to the response header map
		// This lets the user know what request methods are supported for this particular URL
		w.Header().Set("Allow", "POST")

		// 405: Method not allowed
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	// Also catches other paths that dont have a more specific handler 
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")

	// Listen on port 4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
