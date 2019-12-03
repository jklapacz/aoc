package day03_test

import (
	"github.com/jklapacz/aoc/day03"
	"strings"
	"testing"
)

type PuzzleTest struct {
	input string
	solution string
}

func TestSolve(t *testing.T) {
	testFirst := `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	solnFirst := 159
	actualFirst := day03.Solve(strings.NewReader(testFirst))
	if solnFirst != actualFirst {
		t.Errorf("solution is incorrect! expected: %v actual %v", solnFirst, actualFirst)
	}
}

