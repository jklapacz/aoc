package main

import (
	"fmt"
	"os"
)

type Puzzle struct {
	filename string
	openPuzzle *os.File
}

func main() {
	fileArg := os.Args[1:]
	if len(fileArg) < 1 {
		panic("No argument given!")
	}
	puzzle := &Puzzle{filename: fileArg[0]}
	cleanup := puzzle.se()
	defer cleanup()
}

func (p Puzzle) setupPuzzle() func() {
	if p.filename == "" {
		panic("no file given!")
	}
	file, err := os.Open(p.filename)
	if err != nil {
		panic(fmt.Sprintf("could not open file!: %v", err))
	}
	p.openPuzzle = file
	return func() { file.Close() }
}

