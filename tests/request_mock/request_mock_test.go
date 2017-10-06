package main

import (
	"net/http"
	"fmt"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>Going Go programming!</title>
		<description>Golang: </description>
	</channel
</rss>
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the neeed to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the GET call", ballotX, err)
			}

			t.Log("\t\tShould be able to make the GET call", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}

			t.Logf("\t\tShould receive a \"%d\" status, %v", statusCode, checkMark)
		}
	}
}
