package main

import (
	"github.com/doppelganger113/goplay/tests/request_mock_2/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Register()

	log.Println("Server listening on port :4000")
	http.ListenAndServe(":4000", nil)
}
