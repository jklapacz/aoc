package day10_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Grid struct {
	topLeft, bottomRight point
}

func iterateThroughGrid(topLeft, bottomRight point, apply func(x, y int)) {
	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			apply(x, y)
		}
	}
}

func (g *Grid) enumerate() *pointSet {
	ps := &pointSet{topLeft: g.topLeft, bottomRight: g.bottomRight}
	apply := func(x, y int) {
		ps.add(point{x, y})
	}
	iterateThroughGrid(g.topLeft, g.bottomRight, apply)
	return ps
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

type pointSet struct {
	topLeft, bottomRight, origin point
	members                      map[point]bool
}

func (ps *pointSet) remove(p point) {
	delete(ps.members, p)
}

func (ps *pointSet) add(p point) {
	if ps.members == nil {
		ps.members = make(map[point]bool, 0)
	}
	ps.members[p] = true
}

func (ps *pointSet) contains(p point) bool {
	_, ok := ps.members[p]
	return ok
}

func pointsInLine(origin, topLeft, bottomRight point, m slope) *pointSet {
	isInBounds := func(p point) bool {
		return p.x >= topLeft.x &&
			p.x <= bottomRight.x &&
			p.y >= topLeft.y &&
			p.y <= bottomRight.y
	}
	// move right
	fmt.Printf("moving in the following direction: %v, slope: %v\n", slopeDirection(m).toString(), m)
	points := &pointSet{topLeft: topLeft, bottomRight: bottomRight, origin: origin}
	if m.run == 0 {
		if slopeDirection(m) == dirUp {
			for y := origin.y; y >= topLeft.y; y-- {
				points.add(point{origin.x, y})
			}
		} else {
			for y := origin.y; y <= bottomRight.y; y++ {
				points.add(point{origin.x, y})
			}
		}
		return points
	}

	getY := func(x int) int {
		b := float64(origin.y) + (float64(m.rise) / float64(m.run) * float64(origin.x))
		fmt.Println("b: ", b)
		return int(float64(float64(m.rise)/float64(m.run))*float64(x)) + int(b)
	}

	makePoint := func(x int) point {
		yval := getY(x)
		return point{
			x: x,
			y: yval,
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
		yval := float64(float64(m.rise)/float64(m.run)) * float64(x)
		return float64(int64(yval)) == yval
	}

	if m.run > 0 {
		for x := origin.x; x <= bottomRight.x; x++ {
			if pointPossible(x) {
				addPossiblePoints(x)
			}
		}
	} else {
		fmt.Println("starting at: ", origin)
		for x := origin.x; x >= topLeft.x; x-- {
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

func slopeDirection(m slope) direction {
	if m.rise == 0 {
		if m.run > 0 {
			return dirRight
		} else if m.run < 0 {
			return dirLeft
		} else {
			return dirNone
		}
	} else if m.rise > 0 {
		if m.run > 0 {
			return dirRightDown
		} else if m.run < 0 {
			return dirLeftDown
		} else {
			return dirDown
		}
	} else {
		if m.run > 0 {
			return dirRightUp
		} else if m.run < 0 {
			return dirLeftUp
		} else {
			return dirUp
		}
	}
}

func (ps *pointSet) boundaries() (point, point) {
	return ps.topLeft, ps.bottomRight
}

func plot(points *pointSet) {
	grid := "===== grid ======\n"
	topLeft, bottomRight := points.boundaries()
	for y := topLeft.y; y <= bottomRight.y; y++ {
		grid += "|"
		for x := topLeft.x; x <= bottomRight.x; x++ {
			if (points.origin == point{x, y}) {
				grid += fmt.Sprintf("* ")
			} else if points.contains(point{x, y}) {
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
	origin := point{4, 3}
	topLeft := point{0, 0}
	bottomRight := point{5, 5}
	m := slope{1, -4}
	points := pointsInLine(origin, topLeft, bottomRight, m)
	plot(points)
	//t.Fatalf("%v", points)
}

func TestTwoPointPlot(t *testing.T) {
	origin := point{4, 3}
	topLeft := point{0, 0}
	bottomRight := point{5, 5}
	target := point{0, 4}
	m := calculateSlope(origin, target)
	assert.Equal(t, slope{1, -4}, m)
	//m := slope{1, -4}
	points := pointsInLine(origin, topLeft, bottomRight, m)
	plot(points)
	assert.Equal(t, 2, len(points.members))
}

func TestEnumeration(t *testing.T) {
	topLeft := point{0, 0}
	bottomRight := point{5, 5}
	g := &Grid{topLeft: topLeft, bottomRight: bottomRight}
	allPoints := g.enumerate().members
	for origin := range allPoints {
		pointsWithoutOrigin := g.enumerate()
		pointsWithoutOrigin.remove(origin)
		for target := range pointsWithoutOrigin.members {
			m := calculateSlope(origin, target)
			alignedPoints := pointsInLine(origin, g.topLeft, g.bottomRight, m)
			//for p := range alignedPoints.members {
			//	pointsWithoutOrigin.remove(p)
			//}
			fmt.Printf("\n\n==============origin: %v, target: %v slope: %v\n", origin, target, m)
			plot(alignedPoints)
		}
	}
}
