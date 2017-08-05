package game

import "testing"

func TestGameRun(t *testing.T) {
	for i := 0; i < 12; i++ {
		g := New(100, 98)
		g.Run()
	}
}
