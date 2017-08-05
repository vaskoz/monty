package game

import (
	"math/rand"
	"time"
)

type Outcome int

const (
	FirstPickWon Outcome = iota
	SecondPickWon
	NoPickWon
)

type Game interface {
	Run() Outcome
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

func (g *game) Run() Outcome {
	prize := g.generator.Intn(g.doors)
	firstPick := g.generator.Intn(g.doors)
	if firstPick == prize {
		return FirstPickWon
	}
	remainingDoors := g.doors
	for i := 0; i < g.reveals; i++ {
		reveal := g.generator.Intn(remainingDoors)
		if reveal == firstPick || reveal == prize {
			i--
			continue
		}
		if reveal < prize {
			prize--
		}
		if reveal < firstPick {
			firstPick--
		}
		remainingDoors--
	}
	secondPick := g.generator.Intn(remainingDoors)
	if secondPick == firstPick {
		if secondPick+1 < remainingDoors {
			secondPick++
		} else {
			secondPick--
		}
	}
	if secondPick == prize {
		return SecondPickWon
	}
	return NoPickWon
}
