package main

import (
	"log"
	"time"
	"github.com/doppelganger113/goplay/concurrency/work"
	"sync"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"json",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

const maxGoroutines = 2

func main() {
	p := work.New(maxGoroutines)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name: name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
