package main

import (
	"net/http"
	"log"
	"flag"
	"goplay/examples/chat_example/chat"
	"goplay/examples/chat_example/trace"
	"os"
)

func main() {

	var addr = flag.String("addr", ":3000", "The addr of the application")
	flag.Parse()
	r := chat.NewRoom()
	r.trace = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	log.Print("Started room")
	go r.run()

	log.Print("Starting webserver on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
