package trace

import (
	"fmt"
	"io"
)

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}

type Tracer interface {
	Trace(...interface{})
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}
