package simulate

import (
	"github.com/vaskoz/monty/game"
)

// Simulator executes multiple Monty-Hall games concurrently
type Simulator interface {
	Run() (int, int, int)
}

type simulator struct {
	gameList []game.Game
}

// New creates a simulator that executes the games concurrently
func New(doors, reveals, games int) Simulator {
	gameList := make([]game.Game, games)
	for i := 0; i < games; i++ {
		gameList[i] = game.New(doors, reveals)
	}
	return &simulator{
		gameList: gameList,
	}
}

func (s *simulator) Run() (int, int, int) {
	result := make(chan game.Outcome, len(s.gameList))
	for _, g := range s.gameList {
		go func(g game.Game) {
			result <- g.Run()
		}(g)
	}
	first, second, none := 0, 0, 0
	for range s.gameList {
		if out := <-result; out == game.FirstPickWon {
			first++
		} else if out == game.SecondPickWon {
			second++
		} else if out == game.NoPickWon {
			none++
		}
	}
	return first, second, none
}
