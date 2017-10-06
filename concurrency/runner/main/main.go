package main

import (
	"github.com/doppelganger113/goplay/concurrency/runner"
	"time"
	"log"
	"os"
)

const timeout = 3 * time.Second

func main() {
	r := runner.New(timeout)
	r.Add(
		createTask(),
		createTask(),
		createTask(),
	)

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor ~ Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
