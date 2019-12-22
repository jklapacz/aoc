package point

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

// Manhattan calculates the manhattan distance between Points a and b
func Manhattan(a, b Point) int {
	distX := int(math.Abs(float64(a.X - b.X)))
	distY := int(math.Abs(float64(a.Y - b.Y)))
	return distX + distY
}

// FindClosest takes an origin and a set of points and returns the shortest distance
func FindClosest(origin Point, candidates ...Point) Point {
	if len(candidates) == 0 {
		return origin
	}
	currentDistance := math.MaxInt64
	closestPoint := Point{math.MaxInt64, math.MaxInt64}
	for _, potential := range candidates {
		if potential == origin {
			continue
		}
		distance := Manhattan(origin, potential)
		if distance < currentDistance {
			currentDistance = distance
			closestPoint = potential
		}
	}
	return closestPoint
}

type Slope struct {
	Rise, Run int
}

func (s Slope) val() float64 {
	return float64(s.Rise) / float64(s.Run)
}

func CalculateSlope(origin, target Point) Slope {
	return Slope{
		target.Y - origin.Y,
		target.X - origin.X,
	}
}

type Line struct {
	start, end Point
	m          Slope
	b          float64
}

func CreateLine(start, end Point) Line {
	line := Line{start: start, end: end}
	line.m = CalculateSlope(start, end)
	line.b = float64(start.Y) - float64(start.X)*line.m.val()
	return line
}

func (l Line) LineFunc() func(x int) float64 {
	fmt.Printf("y = %vx + %v\n", l.m.val(), l.b)
	return func(x int) float64 {
		yval := float64(x)*l.m.val() + l.b
		fmt.Println("y = ", yval)
		return yval
	}
}

func (l Line) IsValidPoint(x int) bool {
	yval := l.LineFunc()(x)
	return float64(int64(yval)) == yval
}

func (l Line) CreatePoint(x int) Point {
	return Point{
		x,
		int(l.LineFunc()(x)),
	}
}
