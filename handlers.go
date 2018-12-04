package main

import (
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Portfold" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, _ = w.Write([]byte("Hello from Portfold"))
}

func (app *application) help(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Help"))
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("health check"))
}
