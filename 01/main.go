package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	puzzle.Solve()
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

// mostly for debugging, iterate through the file
func (p *Puzzle) dumpContents() {
	if p.instance == nil {
		log.Fatal("puzzle is not instantiated!")
	}
	for p.instance.Scan() {
		fmt.Println(p.instance.Text())
	}
}

type moduleWeight = int
type fuelCost = int

func (p *Puzzle) Solve() {
	fuelSum := 0
	for p.instance.Scan() {
		currentInput := p.instance.Text()
		weight := turnInputIntoWeight(currentInput)
		fuelSum += calculateFuelCost(weight)
	}
	fmt.Println("=== answer: ", fuelSum)
}

func turnInputIntoWeight(input string) moduleWeight {
	weight, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal("invalid module weight input", err)
	}

	return moduleWeight(weight)
}

func calculateFuelCost(m moduleWeight) fuelCost {
	return (m / 3) - 2
}

