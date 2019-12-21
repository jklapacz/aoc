package point

import "math"

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

func Line(start, end Point) func(x int) Point {
	m := CalculateSlope(start, end)
	b := float64(start.X)*m.val() + float64(start.Y)
	lineFunc := func(x int) Point {
		yval := float64(x)*m.val() + b
		return Point{x, int(yval)}
	}
	return lineFunc
}
