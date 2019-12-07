package day06_test

import (
	"fmt"
	"github.com/jklapacz/aoc/day06"
	"strings"
	"testing"
)

func TestOrbitMap_Parse(t *testing.T) {
	om := &day06.OrbitMap{Orbits:make(map[string]*day06.Orbitter)}
	input := "COM)B"
	om.Parse(input)
	for _, orbitter := range om.Orbits {
		fmt.Println(orbitter.ToString())
	}
}

func TestSolve(t *testing.T) {
	input := `COM)B
B)C
C)D`
	expected := 6
	actual := day06.Solve(strings.NewReader(input))
	if expected != actual {
		t.Logf("expected to get %v actually got %v", expected, actual)
		t.Fail()
	}
}

func TestExampleInput(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`
	expected := 42
	actual := day06.Solve(strings.NewReader(input))
	if expected != actual {
		t.Logf("expected to get %v actually got %v", expected, actual)
		t.Fail()
	}
}
