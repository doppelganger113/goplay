package main

import (
	"flag"
	"goplay.com/m/v2/examples/chat_example/chat"
	"goplay.com/m/v2/examples/chat_example/trace"
	"log"
	"net/http"
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
