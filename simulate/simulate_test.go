package simulate

import (
	"os"
	"testing"
)

func TestSimulateRun(t *testing.T) {
	sim := New(3, 1, 1, os.Stdout)
	sim.Run()
}
