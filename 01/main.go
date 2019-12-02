package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Puzzle struct {
	filename string
	instance *PuzzleInstance
}

type PuzzleInstance = bufio.Scanner

func main() {
	fileArg := os.Args[1:]
	if len(fileArg) < 1 {
		panic("No argument given!")
	}
	puzzle := &Puzzle{filename: fileArg[0]}
	cleanup := puzzle.setupPuzzle()
	defer cleanup()
	puzzle.dumpContents()
}

func (p *Puzzle) setupPuzzle() func() {
	if p.filename == "" {
		panic("no file given!")
	}
	file, err := os.Open(p.filename)
	if err != nil {
		panic(fmt.Sprintf("could not open file!: %v", err))
	}
	p.instance = instantiatePuzzle(file)
	return func() { file.Close() }
}

func instantiatePuzzle(file *os.File) *PuzzleInstance {
	return bufio.NewScanner(file)
}

func (p *Puzzle) dumpContents() {
	if p.instance == nil {
		log.Fatal("puzzle is not instantiated!")
	}
	for p.instance.Scan() {
		fmt.Println(p.instance.Text())
	}
}
