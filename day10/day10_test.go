package day10_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Grid struct {
	points [][]int
}

type point struct {
	x, y int
}

type slope struct {
	rise, run int
}

func calculateSlope(origin, target point) slope {
	return slope{
		rise: target.y - origin.y,
		run:  target.x - origin.x,
	}
}

func pointsInLine(origin, topLeft, bottomRight point, m slope) []point {
	isInBounds := func(p point) bool {
		return p.x >= topLeft.x &&
			p.x <= bottomRight.x &&
			p.y >= topLeft.y &&
			p.y <= bottomRight.y
	}
	if m.run == 0 {
		m.run = math.MaxInt64
	}

	makePoint := func(x int) point {
		// vertical line
		yval := float64(float64(m.rise)/float64(m.run)) * float64(x)
		return point{
			x: x,
			y: int(yval),
		}
	}

	var points []point

	addPossiblePoints := func(x int) {
		possiblePoint := makePoint(x)
		if isInBounds(possiblePoint) {
			points = append(points, possiblePoint)
		}
	}

	pointPossible := func(x int) bool {
		yval := float64(float64(m.rise)/float64(m.run)) * float64(x)
		return float64(int64(yval)) == yval
	}

	// move right
	for x := origin.x; x <= bottomRight.x; x++ {
		if pointPossible(x) {
			addPossiblePoints(x)
		}
	}

	for x := origin.x - 1; x >= topLeft.x; x-- {
		if pointPossible(x) {
			addPossiblePoints(x)
		}
	}

	return points
}

func TestPointLine(t *testing.T) {
	origin := point{2, 5}
	topLeft := point{0, 0}
	bottomRight := point{10, 10}
	m := slope{3, 2}
	t.Fatalf("%v", pointsInLine(origin, topLeft, bottomRight, m))
}

func TestSlope(t *testing.T) {
	type scenario struct {
		start, end    point
		expectedSlope slope
	}
	scenarios := []scenario{
		{point{0, 0}, point{3, 1}, slope{1, 3}},
		{point{0, 0}, point{-3, 1}, slope{1, -3}},
	}
	for _, s := range scenarios {
		assert.Equal(t, s.expectedSlope, calculateSlope(s.start, s.end))
	}
}
