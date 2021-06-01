package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value to small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared resource")
		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil
	default:
		log.Println("Acquired:", "New resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.Lock()
	defer p.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "In Queue")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.Lock()
	defer p.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
