package interactions

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
}

type engine struct {
	speed int
}

func NewEngine() Engine {
	return &engine{}
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	e.speed += 10
}

func (e *engine) Decel() {
	e.speed -= 10
}
