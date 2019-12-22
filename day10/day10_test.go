package day10_test

import (
	"fmt"
	"testing"

	"github.com/gookit/color"
	point "github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

type Grid struct {
	topLeft, bottomRight point.Point
}

func iterateThroughGrid(topLeft, bottomRight point.Point, apply func(x, y int)) {
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			apply(x, y)
		}
	}
}

func (g *Grid) enumerate() *pointSet {
	ps := &pointSet{topLeft: g.topLeft, bottomRight: g.bottomRight}
	apply := func(x, y int) {
		ps.add(point.Point{X: x, Y: y})
	}
	iterateThroughGrid(g.topLeft, g.bottomRight, apply)
	return ps
}

type pointSet struct {
	topLeft, bottomRight, origin point.Point
	members                      map[point.Point]bool
}

func (ps *pointSet) remove(p point.Point) {
	delete(ps.members, p)
}

func (ps *pointSet) add(p point.Point) {
	if ps.members == nil {
		ps.members = make(map[point.Point]bool, 0)
	}
	ps.members[p] = true
}

func (ps *pointSet) contains(p point.Point) bool {
	_, ok := ps.members[p]
	return ok
}

func pointsInLine(origin, target, topLeft, bottomRight point.Point) *pointSet {
	isInBounds := func(p point.Point) bool {
		return p.X >= topLeft.X &&
			p.X <= bottomRight.X &&
			p.Y >= topLeft.Y &&
			p.Y <= bottomRight.Y
	}
	m := point.CalculateSlope(origin, target)
	line := point.CreateLine(origin, target)
	// move right
	fmt.Printf("moving in the following direction: %v, slope: %v\n", slopeDirection(m).toString(), m)
	points := &pointSet{topLeft: topLeft, bottomRight: bottomRight, origin: origin}
	if m.Run == 0 {
		if slopeDirection(m) == dirUp {
			for y := origin.Y; y >= topLeft.Y; y-- {
				points.add(point.Point{X: origin.X, Y: y})
			}
		} else {
			for y := origin.Y; y <= bottomRight.Y; y++ {
				points.add(point.Point{X: origin.X, Y: y})
			}
		}
		return points
	}

	addPossiblePoints := func(x int) {
		possiblePoint := line.CreatePoint(x)
		// fmt.Println("possible: ", possiblePoint)
		if isInBounds(possiblePoint) {
			// fmt.Println("in bounds!")
			points.add(possiblePoint)
		}
	}

	if m.Run > 0 {
		for x := origin.X; x <= bottomRight.X; x++ {
			if line.IsValidPoint(x) {
				addPossiblePoints(x)
			}
		}
	} else {
		for x := origin.X; x >= topLeft.X; x-- {
			if line.IsValidPoint(x) {
				addPossiblePoints(x)
			}
		}
	}

	return points
}

type direction int

const (
	dirRight direction = iota
	dirRightDown
	dirDown
	dirLeftDown
	dirLeft
	dirLeftUp
	dirUp
	dirRightUp
	dirNone
)

func (d direction) toString() string {
	switch d {
	case dirRight:
		return "right"
	case dirRightDown:
		return "right & down"
	case dirDown:
		return "down"
	case dirLeftDown:
		return "left & down"
	case dirLeft:
		return "left"
	case dirLeftUp:
		return "left & up"
	case dirUp:
		return "up"
	case dirRightUp:
		return "right & up"
	default:
		return "standstill"
	}
}

func slopeDirection(m point.Slope) direction {
	if m.Rise == 0 {
		if m.Run > 0 {
			return dirRight
		} else if m.Run < 0 {
			return dirLeft
		} else {
			return dirNone
		}
	} else if m.Rise > 0 {
		if m.Run > 0 {
			return dirRightDown
		} else if m.Run < 0 {
			return dirLeftDown
		} else {
			return dirDown
		}
	} else {
		if m.Run > 0 {
			return dirRightUp
		} else if m.Run < 0 {
			return dirLeftUp
		} else {
			return dirUp
		}
	}
}

func (ps *pointSet) boundaries() (point.Point, point.Point) {
	return ps.topLeft, ps.bottomRight
}

func plot(points *pointSet) {
	grid := "===== grid ======\n"
	topLeft, bottomRight := points.boundaries()
	red := color.FgRed.Render
	blue := color.FgBlue.Render
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		grid += "|"
		for x := topLeft.X; x <= bottomRight.X; x++ {
			if (points.origin == point.Point{x, y}) {
				grid += fmt.Sprintf("%s ", blue("+"))
			} else if points.contains(point.Point{x, y}) {
				grid += fmt.Sprintf("%s ", red("+"))
			} else {
				grid += fmt.Sprintf(". ")
			}
		}
		grid += "|\n"
	}
	fmt.Println(grid)
}

func TestPointLine(t *testing.T) {
	origin := point.Point{5, 3}
	target := point.Point{2, 4}
	topLeft := point.Point{0, 0}
	bottomRight := point.Point{5, 5}
	points := pointsInLine(origin, target, topLeft, bottomRight)
	plot(points)
	t.Fatalf("%v", points)
}

func TestTwoPointPlot(t *testing.T) {
	origin := point.Point{5, 3}
	topLeft := point.Point{0, 0}
	bottomRight := point.Point{5, 5}
	target := point.Point{2, 4}
	points := pointsInLine(origin, target, topLeft, bottomRight)
	plot(points)
	assert.Equal(t, 2, len(points.members))
}

func TestEnumeration(t *testing.T) {
	topLeft := point.Point{0, 0}
	bottomRight := point.Point{5, 5}
	g := &Grid{topLeft: topLeft, bottomRight: bottomRight}
	allPoints := g.enumerate().members
	for origin := range allPoints {
		pointsWithoutOrigin := g.enumerate()
		pointsWithoutOrigin.remove(origin)
		for target := range pointsWithoutOrigin.members {
			alignedPoints := pointsInLine(origin, target, g.topLeft, g.bottomRight)
			fmt.Printf("\n\n==============origin: %v, target: %v \n", origin, target)
			plot(alignedPoints)
		}
	}
}
