package graph_test

import (
	"testing"

	point "github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestSlope(t *testing.T) {
	type scenario struct {
		start, end    point.Point
		expectedSlope point.Slope
	}
	scenarios := []scenario{
		{point.Point{0, 0}, point.Point{3, 1}, point.Slope{1, 3}},
		{point.Point{0, 0}, point.Point{-3, 1}, point.Slope{1, -3}},
		{point.Point{0, 0}, point.Point{3, 0}, point.Slope{0, 3}},
		{point.Point{0, 0}, point.Point{0, 3}, point.Slope{3, 0}},
	}
	for _, s := range scenarios {
		assert.Equal(t, s.expectedSlope, point.CalculateSlope(s.start, s.end))
	}
}
