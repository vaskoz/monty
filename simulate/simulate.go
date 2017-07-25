package simulate

import (
	"io"
	"sync"

	"github.com/vaskoz/monty/game"
)

type Simulator interface {
	Run()
}

type simulator struct {
	writer   io.Writer
	gameList []game.Game
}

func New(doors, reveals, games int, writer io.Writer) Simulator {
	gameList := make([]game.Game, games)
	for i := 0; i < games; i++ {
		gameList[i] = game.New(doors, reveals)
	}
	return &simulator{
		gameList: gameList,
		writer:   writer,
	}
}

func (s *simulator) Run() {
	var wg sync.WaitGroup
	wg.Add(len(s.gameList))
	for _, g := range s.gameList {
		go func(g game.Game) {
			g.Run()
			wg.Done()
		}(g)
	}
	wg.Wait()
}
