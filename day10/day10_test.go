package day10_test

import (
	"fmt"
	"testing"

	"github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestPointLine(t *testing.T) {
	origin := graph.Point{X: 5, Y: 3}
	target := graph.Point{X: 2, Y: 4}
	topLeft := graph.Point{X: 0, Y: 0}
	bottomRight := graph.Point{X: 5, Y: 5}
	points := graph.PointsInLine(origin, target, topLeft, bottomRight)
	graph.Plot(points)
}

func TestTwoPointPlot(t *testing.T) {
	origin := graph.Point{X: 5, Y: 3}
	topLeft := graph.Point{X: 0, Y: 0}
	bottomRight := graph.Point{X: 5, Y: 5}
	target := graph.Point{X: 2, Y: 4}
	points := graph.PointsInLine(origin, target, topLeft, bottomRight)
	graph.Plot(points)
	assert.Equal(t, 2, len(points.Members))
}

func TestEnumeration(t *testing.T) {
	topLeft := graph.Point{X: 0, Y: 0}
	bottomRight := graph.Point{X: 5, Y: 5}
	g := &graph.Grid{TopLeft: topLeft, BottomRight: bottomRight}
	allPoints := g.Enumerate().Members
	for origin := range allPoints {
		pointsWithoutOrigin := g.Enumerate()
		pointsWithoutOrigin.Remove(origin)
		for target := range pointsWithoutOrigin.Members {
			alignedPoints := graph.PointsInLine(origin, target, g.TopLeft, g.BottomRight)
			fmt.Printf("\n\n==============origin: %v, target: %v \n", origin, target)
			graph.Plot(alignedPoints)
		}
	}
}
