// Package handlers provides the endpoints for the web servers
package handlers

import (
	"net/http"
	"encoding/json"
)

// Registers the routes with the web server
func Register() {
	http.HandleFunc("/people", People)
}

// Returns a single person json document
func People(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{"bill", "bill@gmail.com"}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(&u)
}
