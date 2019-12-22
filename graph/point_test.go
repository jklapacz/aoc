package graph_test

import (
	"testing"

	point "github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestManhattan(t *testing.T) {
	pointA := point.Point{0, 0}
	pointB := point.Point{0, 0}
	distance := point.Manhattan(pointA, pointB)
	assert.Equal(t, 0, distance)

	pointB = point.Point{5, 0}
	distance = point.Manhattan(pointA, pointB)
	assert.Equal(t, 5, distance)

	pointB = point.Point{5, 5}
	distance = point.Manhattan(pointA, pointB)
	assert.Equal(t, 10, distance)
}

func TestFindClosest(t *testing.T) {
	origin := point.Point{5, 5}
	closestPoint := point.FindClosest(origin,
		point.Point{5, 7},
		point.Point{5, 9},
		point.Point{5, 4},
	)

	assert.Equal(t, point.Point{5, 4}, closestPoint)
}

func TestLine(t *testing.T) {
	start := point.Point{0, 0}
	end := point.Point{3, 3}
	line := point.CreateLine(start, end)
	y := line.LineFunc()(2)
	assert.Equal(t, 2.0, y)

	start = point.Point{0, 9}
	end = point.Point{1, 4}
	line = point.CreateLine(start, end)
	y = line.LineFunc()(2)
	assert.Equal(t, -1.0, y)

	start = point.Point{2, 2}
	end = point.Point{3, 3}
	line = point.CreateLine(start, end)
	y = line.LineFunc()(0)
	assert.Equal(t, 0., y)
}
