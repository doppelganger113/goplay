package interactions

type Train struct {
	engine Engine
}

func NewTrain() *Train {
	return &Train{
		engine: NewEngine(),
	}
}

func (t *Train) Go() {
	t.engine.Accel()
	t.engine.Accel()
}

func (t *Train) Stop() {
	t.engine.Decel()
	t.engine.Decel()
}

func (t *Train) Speed() int {
	return t.engine.Speed()
}
