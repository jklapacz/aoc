package main

import (
	"fmt"
	"github.com/jklapacz/aoc"
	"log"
	"os"
	"strconv"
)

type Puzzle struct {
	*aoc.Puzzle
}

func main() {
	fileArg := os.Args[1:]
	if len(fileArg) < 1 {
		panic("No argument given!")
	}
	puzzle := &Puzzle{&aoc.Puzzle{Filename: fileArg[0]}}
	cleanup := puzzle.SetupPuzzle()
	defer cleanup()

	puzzle.Solve()
}


type moduleWeight = int
type fuelCost = int

func (p *Puzzle) Solve() {
	fuelSum := 0
	for p.Instance.Scan() {
		currentInput := p.Instance.Text()
		weight := turnInputIntoWeight(currentInput)
		fuelSum += calculateFuelCostCost(weight)
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

func calculateFuelCostCost(m moduleWeight) fuelCost {
	calculatedCost := calculateFuelCost(m)
	if calculatedCost <= 0 {
		return 0
	}

	return calculatedCost + calculateFuelCostCost(calculatedCost)
}

