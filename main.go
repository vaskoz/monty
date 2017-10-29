package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/vaskoz/monty/simulate"
)

const loggerPrefix = "monty-"

var stdout io.Writer = os.Stdout
var logger = log.New(stdout, loggerPrefix, log.LstdFlags)
var exit = os.Exit

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
	simulator := simulate.New(*doors, *reveals, *games)
	first, second, none := simulator.Run()
	logger.Println("First choice wins:", first)
	logger.Println("Second choice wins:", second)
	logger.Println("Didn't win:", none)
}
