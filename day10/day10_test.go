package day10_test

import (
	"fmt"
	"testing"

	"github.com/jklapacz/aoc/point"
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
		ps.add(point.Point{x, y})
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

func pointsInLine(origin, topLeft, bottomRight point.Point, m point.Slope) *pointSet {
	isInBounds := func(p point.Point) bool {
		return p.X >= topLeft.X &&
			p.X <= bottomRight.X &&
			p.Y >= topLeft.Y &&
			p.Y <= bottomRight.Y
	}
	// move right
	fmt.Printf("moving in the following direction: %v, slope: %v\n", slopeDirection(m).toString(), m)
	points := &pointSet{topLeft: topLeft, bottomRight: bottomRight, origin: origin}
	if m.Run == 0 {
		if slopeDirection(m) == dirUp {
			for y := origin.Y; y >= topLeft.Y; y-- {
				points.add(point.Point{origin.X, y})
			}
		} else {
			for y := origin.Y; y <= bottomRight.Y; y++ {
				points.add(point.Point{origin.X, y})
			}
		}
		return points
	}

	getY := func(x int) int {
		b := float64(origin.Y) + (float64(m.Rise) / float64(m.Run) * float64(origin.X))
		fmt.Println("b: ", b)
		return int(float64(float64(m.Rise)/float64(m.Run))*float64(x)) + int(b)
	}

	makePoint := func(x int) point.Point {
		yval := getY(x)
		return point.Point{
			X: x,
			Y: yval,
		}
	}

	addPossiblePoints := func(x int) {
		possiblePoint := makePoint(x)
		if isInBounds(possiblePoint) {
			fmt.Println("in bounds!")
			points.add(possiblePoint)
		}
	}

	pointPossible := func(x int) bool {
		yval := float64(float64(m.Rise)/float64(m.Run)) * float64(x)
		return float64(int64(yval)) == yval
	}

	if m.Run > 0 {
		for x := origin.X; x <= bottomRight.X; x++ {
			if pointPossible(x) {
				addPossiblePoints(x)
			}
		}
	} else {
		fmt.Println("starting at: ", origin)
		for x := origin.X; x >= topLeft.X; x-- {
			fmt.Printf("considering: %v, %v\n", x, getY(x))
			if pointPossible(x) {
				fmt.Printf("adding!: %v, %v\n", x, getY(x))
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
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		grid += "|"
		for x := topLeft.X; x <= bottomRight.X; x++ {
			if (points.origin == point.Point{x, y}) {
				grid += fmt.Sprintf("* ")
			} else if points.contains(point.Point{x, y}) {
				grid += fmt.Sprintf("+ ")
			} else {
				grid += fmt.Sprintf(". ")
			}
		}
		grid += "|\n"
	}
	fmt.Println(grid)
}

func TestPointLine(t *testing.T) {
	origin := point.Point{4, 3}
	topLeft := point.Point{0, 0}
	bottomRight := point.Point{5, 5}
	m := point.Slope{1, -4}
	points := pointsInLine(origin, topLeft, bottomRight, m)
	plot(points)
	//t.Fatalf("%v", points)
}

func TestTwoPointPlot(t *testing.T) {
	origin := point.Point{4, 3}
	topLeft := point.Point{0, 0}
	bottomRight := point.Point{5, 5}
	target := point.Point{0, 4}
	m := point.CalculateSlope(origin, target)
	assert.Equal(t, point.Slope{1, -4}, m)
	//m := slope{1, -4}
	points := pointsInLine(origin, topLeft, bottomRight, m)
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
			m := point.CalculateSlope(origin, target)
			alignedPoints := pointsInLine(origin, g.topLeft, g.bottomRight, m)
			//for p := range alignedPoints.members {
			//	pointsWithoutOrigin.remove(p)
			//}
			fmt.Printf("\n\n==============origin: %v, target: %v slope: %v\n", origin, target, m)
			plot(alignedPoints)
		}
	}
}
