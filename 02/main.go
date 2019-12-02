package main

import (
	"fmt"
	"github.com/jklapacz/aoc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("argument needed")
	}
	filename := os.Args[1]
	puzzle := &aoc.Puzzle{Filename:filename}
	cleanup := puzzle.SetupPuzzle()
	defer cleanup()
	solve(puzzle)
}

func solve(p *aoc.Puzzle) {
	fmt.Println("Solving!")
	p.DumpContents()
}