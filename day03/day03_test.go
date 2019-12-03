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
	type scenario struct {
		input string
		output int
	}
	scenarios := []scenario{
		{
			`R8,U5,L5,D3
U7,R6,D4,L4`,
6,
		},
		{
			`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			159,
		},
		{
			`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			135,
		},
		{
			`L10,U8
U8,L10,U7`,
			18,
		},
	}
	for _, s := range scenarios {
		actual := day03.Solve(strings.NewReader(s.input))
		if actual != s.output {
			t.Errorf("solution is incorrect! expected: %v actual %v", s.output, actual)
		}
	}
}

func TestIntersection(t *testing.T) {
	type scenario struct {
		segmentA day03.Segment
		segmentB day03.Segment
		intersection *day03.Point
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
			intersection: &day03.Point{2, 2},
		},
		{
			segmentA: day03.Segment{
				day03.Point{0,0},
				day03.Point{0 ,4},
			},
			segmentB : day03.Segment{
				day03.Point{0, 2},
				day03.Point{0, 6},
			},
			intersection: &day03.Point{0, 2},
		},
		{
			segmentA: day03.Segment{
				day03.Point{0,0},
				day03.Point{4 ,0},
			},
			segmentB : day03.Segment{
				day03.Point{2, 0},
				day03.Point{6, 0},
			},
			intersection: &day03.Point{2, 0},
		},
		{
			segmentA: day03.Segment{
				day03.Point{3,5},
				day03.Point{3 ,2},
			},
			segmentB : day03.Segment{
				day03.Point{6, 3},
				day03.Point{2, 3},
			},
			intersection: &day03.Point{3, 3},
		},
		{
			segmentA: day03.Segment{
				day03.Point{0,0},
				day03.Point{4 ,0},
			},
			segmentB : day03.Segment{
				day03.Point{2, 0},
				day03.Point{2, 4},
			},
			intersection: &day03.Point{2, 0},
		},
		{
			segmentA: day03.Segment{
				day03.Point{0,0},
				day03.Point{4 ,4},
			},
			segmentB : day03.Segment{
				day03.Point{0, 0},
				day03.Point{4, 4},
			},
			intersection: &day03.Point{1, 1},
		},
		{
			segmentA: day03.Segment{
				day03.Point{0,1},
				day03.Point{ 1,2},
			},
			segmentB : day03.Segment{
				day03.Point{1, 0},
				day03.Point{2, 1},
			},
			intersection: nil,
		},
		{
			segmentA: day03.Segment{
				day03.Point{-10,0},
				day03.Point{ -10,8},
			},
			segmentB : day03.Segment{
				day03.Point{-10, 8},
				day03.Point{-10, 12},
			},
			intersection: &day03.Point{-10, 8},
		},
	}
	for _, s := range scenarios {
		t.Log("Testing: ", s.segmentA, s.segmentB)
		solution := day03.Intersection(s.segmentA, s.segmentB)
		if s.intersection == nil {
			if len(solution) != 0 {
				t.Fatalf(" bad length expected: %v actual :%v", s.intersection, *solution[0])
			}
			//soln := solution[0]
			//if soln != nil {
			//	t.Fail()
			//}
		} else {
			if len(solution) == 0 {
				t.Fatalf(" bad length expected: %v actual :%v", s.intersection, solution)
			}
			soln := solution[0]
			if soln != nil && *soln != *s.intersection {
				t.Fatalf("invalid input expected: %v actual :%v", *s.intersection, *soln)
			}
		}
	}
}