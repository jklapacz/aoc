package main

import (
	"fmt"
	"github.com/jklapacz/aoc/day03"
	"github.com/jklapacz/aoc/day04"
	"github.com/jklapacz/aoc/day06"
	"github.com/jklapacz/aoc/puzzle"
)

func main() {
	//solveDay03()
	//solveDay04()
	solveDay06()
}

func solveDay03() {
	filename := "foobar.txt"
	currentPuzzle := &puzzle.Puzzle{Filename:filename}
	cleanup := currentPuzzle.Setup()
	defer cleanup()
	//fmt.Println("Distance: ", day03.Solve(currentPuzzle))
	fmt.Println("Cost : ", day03.SolvePart2(currentPuzzle))
	//fmt.Println("Cost : ", day03.SolvePart2(currentPuzzle))
}

func solveDay04() {
	start := 128392
	end := 643281
	fmt.Println("Total: ", day04.TotalValidPasscodesInRange(start, end))
}

func solveDay06() {
	filename := "day06/input.txt"
	currentPuzzle := &puzzle.Puzzle{Filename:filename}
	cleanup := currentPuzzle.Setup()
	defer cleanup()
	fmt.Println("Orbits: ", day06.Solve(currentPuzzle))
}
