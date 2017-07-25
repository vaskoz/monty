package game

import (
	"math/rand"
	"time"
)

type Game interface {
	Run()
}

type game struct {
	doors, reveals int
	generator      *rand.Rand
}

func New(doors, reveals int) Game {
	return &game{
		doors:     doors,
		reveals:   reveals,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *game) Run() {
}
