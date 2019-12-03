package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	DirUp = "U"
	DirDown = "D"
	DirRight = "R"
	DirLeft = "L"
)

type WirePath = []Point

func Solve(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	wires := make([]WirePath, 2)
	count := 0
	for scanner.Scan() {
		// get the index, every 2 lines are a pair of wires
		wireIdx := int(math.Mod(float64(count), float64(2)))
		wirePath := scanner.Text()
		wires[wireIdx] = stringToPath(wirePath)
		fmt.Println("Line ", count, wirePath, wires[wireIdx])

		if wireIdx == 1 {
			// we can check for wire crossings
			wireCrossings := getCrossings(wires...)
			for _, crossing := range wireCrossings {
				fmt.Println("Crossings!!!", *crossing)
			}
			return FindClosest(Point{0,0}, wireCrossings[1:]...)
		}
		count++
	}
	return 1
}

func getCrossings(wires ...WirePath) []*Point {
	fmt.Println("Checking where wires cross!")
	if len(wires) < 2 {
		return nil
	}
	var crossings []*Point
	truthWire := wires[0]
	truthPrev := Point{0,0}
	wirePrev := Point{0,0}
	for _, wirePoint := range wires[1] {
		wireSegment := Segment{wirePrev, wirePoint}
		for _, truePoint := range truthWire {
			truthSegment := Segment{truthPrev, truePoint}
			fmt.Println("Current Wire: ", wireSegment, "Current truth: ", truthSegment)
			intersectionPoint := Intersection(wireSegment, truthSegment)
			if intersectionPoint != nil {
				fmt.Println("Intersection point: ", intersectionPoint)
				crossings = append(crossings, intersectionPoint)
			}
			truthPrev = truePoint
		}
		wirePrev = wirePoint
	}
	return crossings
}

type Segment struct {
	Start, End Point
}

func Intersection(a, b Segment) *Point {
	x1 := float64(a.Start.X)
	x2 := float64(a.End.X)
	x3 := float64(b.Start.X)
	x4 := float64(b.End.X)
	y1 := float64(a.Start.Y)
	y2 := float64(a.End.Y)
	y3 := float64(b.Start.Y)
	y4 := float64(b.End.Y)

	interval1 := []float64{
		math.Min(x1, x2),
		math.Max(x1, x2),
	}
	interval2 := []float64{
		math.Min(x3, x4),
		math.Max(x3, x4),
	}
	intervaly1 := []float64{
		math.Min(y1, y2),
		math.Max(y1, y2),
	}
	intervaly2 := []float64{
		math.Min(y3, y4),
		math.Max(y3, y4),
	}

	a1 := (y2 - y1)
	b1 := (x1 - x2)
	c1 := a1*x1 + b1*y1

	a2 := (y4 - y3)
	b2 := (x3 - x4)
	c2 := a2*x3 + b2*y3

	det := a1*b2 - a2*b1

	isWithin := func(xCoord, yCoord float64) bool {
		return xCoord >= interval1[0] && xCoord <= interval1[1] &&
			xCoord >= interval2[0] && xCoord <= interval2[1] &&
			yCoord >= intervaly1[0] && yCoord <= intervaly1[1] &&
			yCoord >= intervaly2[0] && yCoord <= intervaly2[1]
	}

	if det == 0 {
		// parallel lines
		return nil
	}

	x := (b2*c1 - b1*c2)/det
	y := (a1*c2 - a2*c1)/det

	if isWithin(x, y) {
		return &Point{int(x), int(y)}
	}
	return nil
}

func stringToPath(input string) WirePath {
	movements := strings.Split(input, ",")
	points := make(WirePath, len(movements))
	currentPoint := Point{0, 0}
	for pointIdx, movement := range movements {
		if len(movement) < 2 {
			log.Fatal("Invalid coordinate instruction", "idx", pointIdx, "instruction", movement)
		}
		delta, err := strconv.ParseInt(movement[1:], 10, 64)
		if err != nil {
			log.Fatal("coordinate is invalid!")
		}
		switch string(movement[0]) {
		case DirUp:
			points[pointIdx] = Point{currentPoint.X, currentPoint.Y + int(delta)}
		case DirDown:
			points[pointIdx] = Point{currentPoint.X, currentPoint.Y - int(delta)}
		case DirLeft:
			points[pointIdx] = Point{currentPoint.X - int(delta), currentPoint.Y}
		case DirRight:
			points[pointIdx] = Point{currentPoint.X + int(delta), currentPoint.Y}
		default:
			log.Fatal("Invalid direction!")
		}
		currentPoint = points[pointIdx]
	}
	return points
}

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
func FindClosest(origin Point, candidates ...*Point) int {
	if len(candidates) == 0 {
		return -1
	}
	currentDistance := int(^uint(0) >> 1)
	for _, potential := range candidates {
		distance := Manhattan(origin, *potential)
		if distance < currentDistance {
			currentDistance = distance
		}
	}
	return currentDistance
}
