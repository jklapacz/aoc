package main

import (
	"github.com/jklapacz/aoc/day03"
	"github.com/jklapacz/aoc/puzzle"
)

func main() {
	solveDay03()
}

func solveDay03() {
	filename := "foobar.txt"
	currentPuzzle := &puzzle.Puzzle{Filename:filename}
	cleanup := currentPuzzle.Setup()
	defer cleanup()
	day03.Solve(currentPuzzle)
}