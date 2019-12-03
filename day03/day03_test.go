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

func TestManhattan(t *testing.T) {
	a := day03.Point{0, 0}
	b := day03.Point{0, 0}
	dist := day03.Manhattan(a, b)
	if dist != 0 {
		t.Error("Distance calculated incorrectly!")
	}
}

func TestFindClosest(t *testing.T) {
	origin := day03.Point{0,0}
	candidates := []*day03.Point{
		{3,3},
		{6,6},
	}
	closestDistance := day03.FindClosest(origin, candidates...)
	if closestDistance != 6 {
		t.Errorf("did not find closest candidate! expected: %v actual: %v\n", 6, closestDistance)
	}
}

func TestSolve(t *testing.T) {
	//testFirst := `R8,U5,L5,D3
//U7,R6,D4,L4`
	testFirst := `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
	solnFirst := 135
	actualFirst := day03.Solve(strings.NewReader(testFirst))
	if solnFirst != actualFirst {
		t.Errorf("solution is incorrect! expected: %v actual %v", solnFirst, actualFirst)
	}
}

func TestIntersection(t *testing.T) {
	type scenario struct {
		segmentA day03.Segment
		segmentB day03.Segment
		intersection day03.Point
	}
	scenarios := []scenario{
		{
			segmentA: day03.Segment{
			day03.Point{0,0},
			day03.Point{4 ,4},
			},
			segmentB : day03.Segment{
			day03.Point{0, 4},
			day03.Point{4, 0},
			},
			intersection: day03.Point{2, 2},
		},
		//{
		//	segmentA: day03.Segment{
		//		day03.Point{0,0},
		//		day03.Point{4 ,4},
		//	},
		//	segmentB : day03.Segment{
		//		day03.Point{0, 0},
		//		day03.Point{2, 2},
		//	},
		//	intersection: day03.Point{0, 0},
		//},
		{
			segmentA: day03.Segment{
				day03.Point{3,5},
				day03.Point{3 ,2},
			},
			segmentB : day03.Segment{
				day03.Point{6, 3},
				day03.Point{2, 3},
			},
			intersection: day03.Point{3, 3},
		},
	}
	for _, s := range scenarios {
		solution := day03.Intersection(s.segmentA, s.segmentB)
		if solution == nil {
			t.Fail()
		}
		if *solution != s.intersection {
			t.Errorf("expected: %v actual :%v", s.intersection, solution)
		}
	}
}