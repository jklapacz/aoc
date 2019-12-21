package point_test

import (
	"testing"

	"github.com/jklapacz/aoc/point"
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

func TestSlope(t *testing.T) {
	type scenario struct {
		start, end    point.Point
		expectedSlope point.Slope
	}
	scenarios := []scenario{
		{point.Point{0, 0}, point.Point{3, 1}, point.Slope{1, 3}},
		{point.Point{0, 0}, point.Point{-3, 1}, point.Slope{1, -3}},
	}
	for _, s := range scenarios {
		assert.Equal(t, s.expectedSlope, point.CalculateSlope(s.start, s.end))
	}
}
