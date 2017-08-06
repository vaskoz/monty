package simulate

import (
	"testing"
)

func TestSimulateRun(t *testing.T) {
	games := 10000
	sim := New(4, 1, games)
	sim.Run()
}
