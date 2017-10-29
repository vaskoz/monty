package game

import (
	"math/rand"
	"time"
)

// Outcome enumerates the possible simulation results
type Outcome int

const (
	// FirstPickWon is the outcome if the first choice was the correct choice
	FirstPickWon Outcome = iota
	// SecondPickWon is the outcome if switching won the game
	SecondPickWon
	// NoPickWon is if neither the first pick nor the second pick would have won
	NoPickWon
)

// Game is how a client runs a simulation of the Monty-Hall problem
type Game interface {
	Run() Outcome
}

type game struct {
	doors, reveals int
	generator      *rand.Rand
}

// New constructs a Monty-Hall simulation
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
