package simulate

import (
	"testing"
)

func TestSimulateRun(t *testing.T) {
	games := 8000
	sim := New(4, 1, games)
	sim.Run()
}
