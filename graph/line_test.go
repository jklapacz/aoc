package graph_test

import (
	"testing"

	"github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestPointsInLine(t *testing.T) {
	type scenario struct {
		start, end graph.Point
		dimensions []graph.Point
		expected   []graph.Point
	}
	scenarios := []scenario{
		{
			start:      graph.Point{X: 2, Y: 2},
			end:        graph.Point{X: 4, Y: 3},
			dimensions: []graph.Point{{X: 0, Y: 0}, {X: 5, Y: 5}},
			expected:   []graph.Point{{X: 2, Y: 2}, {X: 4, Y: 3}},
		},
		{
			start:      graph.Point{X: 5, Y: 3},
			end:        graph.Point{X: 2, Y: 4},
			dimensions: []graph.Point{{X: 0, Y: 0}, {X: 5, Y: 5}},
			expected:   []graph.Point{{X: 5, Y: 3}, {X: 2, Y: 4}},
		},
		{
			start:      graph.Point{X: 4, Y: 5},
			end:        graph.Point{X: 1, Y: 4},
			dimensions: []graph.Point{{X: 0, Y: 0}, {X: 5, Y: 5}},
			expected:   []graph.Point{{X: 4, Y: 5}, {X: 1, Y: 4}},
		},
		{
			start:      graph.Point{X: 0, Y: 2},
			end:        graph.Point{X: 0, Y: 4},
			dimensions: []graph.Point{{X: 0, Y: 0}, {X: 5, Y: 5}},
			expected:   []graph.Point{{X: 0, Y: 2}, {X: 0, Y: 3}, {X: 0, Y: 4}, {X: 0, Y: 5}},
		},
		{
			start:      graph.Point{X: 0, Y: 2},
			end:        graph.Point{X: 0, Y: 0},
			dimensions: []graph.Point{{X: 0, Y: 0}, {X: 5, Y: 5}},
			expected:   []graph.Point{{X: 0, Y: 2}, {X: 0, Y: 1}, {X: 0, Y: 1}},
		},
	}

	for _, s := range scenarios {
		points := graph.PointsInLine(
			s.start,
			s.end,
			s.dimensions[0],
			s.dimensions[1],
		)
		graph.Plot(points)
		assert.Equal(t, len(s.expected), len(points.Members))
		for _, expected := range s.expected {
			t.Logf("Checking if %v is within points", expected)
			assert.Equal(t, true, points.Contains(expected))
		}
	}
}

func TestEnumeration(t *testing.T) {
	dimensions := []graph.Point{
		{X: 0, Y: 0},
		{X: 2, Y: 2},
	}

	g := &graph.Grid{TopLeft: dimensions[0], BottomRight: dimensions[1]}

	points := g.Enumerate()
	assert.Equal(t, 9, len(points.Members))

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			assert.Equal(t, true, points.Contains(graph.Point{x, y}))
		}
	}
}

func BenchmarkEnumeration(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dimensions := []graph.Point{{X: 0, Y: 0}, {X: 3, Y: 3}}
		g := &graph.Grid{TopLeft: dimensions[0], BottomRight: dimensions[1]}

		mapFunc := func(p graph.Point) {
			withoutOrigin := g.Enumerate()
			withoutOrigin.Remove(p)
			innerMap := func(t graph.Point) {
				aligned := graph.PointsInLine(p, t, g.TopLeft, g.BottomRight)
				graph.Plot(aligned)
			}
			withoutOrigin.Map(innerMap)
		}

		g.Enumerate().Map(mapFunc)
	}
}
