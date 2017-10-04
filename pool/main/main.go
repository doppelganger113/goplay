package main

import (
	"github.com/doppelganger113/goplay/pool"
	"log"
	"io"
	"sync/atomic"
	"sync"
	"time"
	"math/rand"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

var idCounter int32

type dbConnection struct {
	ID int32
}

func (con *dbConnection) Close() error {
	log.Println("Close: Connection", con.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutting down program.")
	p.Close()
}
