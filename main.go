package main

import (
	"fmt"
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
	//fmt.Println("Distance: ", day03.Solve(currentPuzzle))
	fmt.Println("Cost : ", day03.SolvePart2(currentPuzzle))
}