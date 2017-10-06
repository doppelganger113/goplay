package handlers

import (
	"net/http"
	"encoding/json"
)

func Register() {
	http.HandleFunc("/people", People)
}

func People(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{"bill", "bill@gmail.com"}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(&u)
}
