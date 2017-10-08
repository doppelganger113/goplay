package interactions

type FakeEngine struct {
	AccelCalls int
	DecelCalls int
}

func NewFakeEngine() *FakeEngine {
	return &FakeEngine{}
}

func (e *FakeEngine) Accel() {
	e.AccelCalls += 1
}

func (e *FakeEngine) Decel() {
	e.DecelCalls += 1
}

func (e *FakeEngine) Speed() int {
	return 0
}
