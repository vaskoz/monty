package game

import "testing"

func TestGameRun(t *testing.T) {
	for i := 0; i < 1000; i++ {
		g := New(4, 1)
		g.Run()
	}
}
