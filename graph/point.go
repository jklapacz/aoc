package graph

import (
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Point struct {
	X int
	Y int
}

var trace *log.Logger

func init() {
	var logger io.Writer
	if os.Getenv("DEBUG") == "true" {
		logger = os.Stdout
	} else {
		logger = ioutil.Discard
	}
	trace = log.New(logger, "point: ", log.Ltime|log.Lshortfile)
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
	return func(x int) float64 {
		return float64(x)*l.m.val() + l.b
	}
}

const floatEqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	absDiff := math.Abs(a - b)
	trace.Printf("abs difference: between %v and %v %v\n", a, b, absDiff)
	isClose := absDiff <= floatEqualityThreshold
	trace.Println("are numbers basically the same?", isClose)
	return isClose
}

func (l Line) IsValidPoint(x int) bool {
	yval := l.LineFunc()(x)
	trace.Printf("checking if %v is valid (%v)\n", x, yval)
	trace.Printf("rounded value: %v\n", float64(math.Round(yval)))
	return almostEqual(float64(math.Round(yval)), yval)
}

func (l Line) CreatePoint(x int) Point {
	y := int(math.Round(l.LineFunc()(x)))
	return Point{
		x,
		y,
	}
}

/*
	III   + 270    IV
        . . | . .
    +  	. . | . .
    1   - - 0 - - + 0
    8   . . | . .
    0 II. . | . . I
	       + 90
*/
func CalculateAngle(origin, target Point) float64 {
	dy := float64(target.Y - origin.Y)
	dx := float64(target.X - origin.X)
	switch {
	case dx == 0 && dy < 0:
		return 0
	case dx == 0 && dy > 0:
		return 180
	case dy == 0 && dx > 0:
		return 90
	case dy == 0 && dx < 0:
		return 270
	}
	trace.Printf("%v/%v\n", dy, dx)
	degrees := rad2deg(math.Atan2(dy, dx))
	if degrees < 0 {
		degrees = 360 + degrees
	}
	return degrees
}

func rad2deg(r float64) float64 {
	return r * (180. / math.Pi)
}
