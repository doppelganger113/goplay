package interactions

import "testing"

func expectSpeed(t *testing.T, speeder Speeder, s int) {
	actual := speeder.Speed()
	if actual != s {
		t.Fatal("Expected", s, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
}
