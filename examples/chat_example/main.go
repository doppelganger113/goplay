package main

import (
	"net/http"
	"log"
	"flag"
	"github.com/doppelganger113/goplay/examples/chat_example/chat"
	"github.com/doppelganger113/goplay/examples/chat_example/trace"
	"os"
)

func main() {

	var addr = flag.String("addr", ":3000", "The addr of the application")
	flag.Parse()
	r := chat.NewRoom()
	r.Tracer = trace.New(os.Stdout)

	http.Handle("/", chat.NewTemplateHandler("chat.html"))
	http.Handle("/room", r)

	log.Print("Started room")
	go r.Run()

	log.Print("Starting webserver on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
