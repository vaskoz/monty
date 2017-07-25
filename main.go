package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/vaskoz/monty/simulate"
)

var stdin io.Reader = os.Stdin
var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var logger *log.Logger = log.New(stderr, "monty", log.LstdFlags)
var exit func(int) = os.Exit

func main() {
	doors := flag.Int("doors", 3, "how many doors in the game")
	reveals := flag.Int("reveals", 1, "how many doors are revealed")
	games := flag.Int("games", 1000, "how many games to simulate")
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		logger.Println("exiting")
		exit(1)
	}
	simulator := simulate.New(*doors, *reveals, *games, stdout)
	simulator.Run()
}
