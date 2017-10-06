package handlers_test

import (
	"github.com/doppelganger113/goplay/tests/request_mock_2/handlers"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	handlers.Register()
}

func TestPeople(t *testing.T) {
	t.Log("Given the need to test people")
	{
		req, err := http.NewRequest("GET", "/people", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a new request", ballotX, err)
		}

		t.Log("\tCreated a request", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != http.StatusOK {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}

		t.Log("Received success code 200!", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tResponse decoding failed", ballotX)
		}

		if u.Name != "bill" || u.Email != "bill@gmail.com" {
			t.Fatalf("Got invalid data: %s %s", u.Name, u.Email)
		}
	}
}
