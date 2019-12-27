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

func TestCalculateAngle(t *testing.T) {
	type scenario struct {
		origin, target point.Point
		result         float64
	}
	scenarios := []scenario{
		{
			point.Point{X: 2, Y: 2},
			point.Point{X: 4, Y: 4},
			135.,
		},
		{
			point.Point{X: 2, Y: 2},
			point.Point{X: 1, Y: 3},
			225.,
		},
		{
			point.Point{X: 2, Y: 2},
			point.Point{X: 1, Y: 1},
			315.,
		},
		{
			point.Point{X: 2, Y: 2},
			point.Point{X: 3, Y: 1},
			45.,
		},
	}

	for _, s := range scenarios {
		assert.Equal(t, point.Deg2rad(s.result), point.CalculateAngle(s.origin, s.target))
	}
}

/*
  0 1 2 3 4
0 . . . 2 6
1 . . 1 3 .
2 . . H 4 7
3 . 5 . . .
4 . . . . .
*/

func TestSortingByAngle(t *testing.T) {
	homeBase := point.Point{2, 2}
	points := []point.Point{
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 1, Y: 3},
	}
	//origin := point.Point{0, 0}
	//points := []point.Point{{X: 2, Y: 2}, {X: 1, Y: 1}, {X: -1, Y: 4}, {X: 0, Y: -1}}
	expectedSorted := []point.Point{
		{X: 2, Y: 1},
		{X: 3, Y: 0},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 1, Y: 3},
		{X: 4, Y: 0},
		{X: 4, Y: 2},
	}

	point.SortByAngle(homeBase, points)
	assert.Equal(t, expectedSorted, points)
	//assert.Equal(t, point.Point{X: 0, Y: -1}, points[0])
	//assert.Equal(t, point.Point{X: 1, Y: 1}, points[1])
	//assert.Equal(t, point.Point{X: 2, Y: 2}, points[2])
	//assert.Equal(t, point.Point{X: -1, Y: 4}, points[3])
}
