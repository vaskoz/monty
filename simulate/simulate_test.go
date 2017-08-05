package simulate

import (
	"testing"
)

func TestSimulateRun(t *testing.T) {
	games := 10000
	sim := New(3, 1, games)
	sim.Run()
}
