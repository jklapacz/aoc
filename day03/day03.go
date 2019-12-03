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
	return -1
}
func SolvePart2(input io.Reader) int {
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
			return FindCheapest(Point{0, 0}, wires, wireCrossings[1:]...)
		}
		count++
	}
	return -1
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
			intersectionPoints := Intersection(wireSegment, truthSegment)
			if intersectionPoints != nil && len(intersectionPoints) > 0 {
				for _, i := range intersectionPoints {
					fmt.Println("Intersection point: ", *i)
				}
				crossings = append(crossings, intersectionPoints...)
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

func Intersection(a, b Segment) []*Point {
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
	run1 := (x2 - x1)
	c1 := a1*x1 + b1*y1

	a2 := (y4 - y3)
	b2 := (x3 - x4)
	run2 := (x4 - x3)
	c2 := a2*x3 + b2*y3

	det := a1*b2 - a2*b1

	isWithin := func(xCoord, yCoord float64) bool {
		return xCoord >= interval1[0] && xCoord <= interval1[1] &&
			xCoord >= interval2[0] && xCoord <= interval2[1] &&
			yCoord >= intervaly1[0] && yCoord <= intervaly1[1] &&
			yCoord >= intervaly2[0] && yCoord <= intervaly2[1]
	}

	parallelPoints := func() []*Point {
		// line goes vertical
		var startPoint, endPoint float64
		var points []*Point
		if x1 == x2 && x1 == x3 && x3 == x4 {
			if math.Min(intervaly1[0], intervaly2[0]) == intervaly1[0] {
				// a is lower
				startPoint = intervaly2[0]
			} else {
				startPoint = intervaly1[0]
			}
			if math.Max(intervaly1[1], intervaly2[1]) == intervaly1[1] {
				endPoint = intervaly2[1]
			} else {
				endPoint = intervaly1[1]
			}
			for yIdx := startPoint; yIdx <= endPoint; yIdx++ {
				points = append(points, &Point{int(x1), int(yIdx)})
			}
		} else if y1 == y2 && y1 == y3 && y3 == y4 {
			// moving horizontally
			if math.Min(interval1[0], interval2[0]) == interval1[0] {
				// a is lower
				startPoint = interval2[0]
			} else {
				startPoint = interval1[0]
			}
			if math.Max(interval1[1], interval2[1]) == interval1[1] {
				endPoint = interval2[1]
			} else {
				endPoint = interval1[1]
			}
			for xIdx := startPoint; xIdx <= endPoint; xIdx++ {
				points = append(points, &Point{int(xIdx), int(y1)})
			}
		} else if b1 == b2 && c1 == c2 {
			var startX, startY, endX, endY float64
			if math.Min(interval1[0], interval2[0]) == interval1[0] {
				// a is lower
				startX = interval2[0]
			} else {
				startX = interval1[0]
			}
			if math.Max(interval1[1], interval2[1]) == interval1[1] {
				endX = interval2[1]
			} else {
				endX = interval1[1]
			}
			if math.Min(intervaly1[0], intervaly2[0]) == intervaly1[0] {
				// a is lower
				startY = intervaly2[0]
			} else {
				startY = intervaly1[0]
			}
			if math.Max(intervaly1[1], intervaly2[1]) == intervaly1[1] {
				endY = intervaly2[1]
			} else {
				endY = intervaly1[1]
			}
			fmt.Println("Parallel test: ", run1, run2, startX, startY, endX, endY, a1, a2, b1, b2, c1, c2)
			for xIdx, yIdx := startX, startY; xIdx <= endX && yIdx <= endY; xIdx, yIdx = xIdx+1, yIdx + (a1/run1) {
				if xIdx == 0 && yIdx == 0 {
					continue
				}
				if isWithin(xIdx, yIdx) {
					points = append(points, &Point{int(xIdx), int(yIdx)})
				}
				fmt.Println("now checking", xIdx, yIdx, isWithin(xIdx, yIdx))
			}

			fmt.Println("Parallel test: ", run1, run2, startX, startY, endX, endY, a1, a2, b1, b2, c1, c2)
			fmt.Printf("oops, b is: %v slope is: %v and %v, x1 is %v x2 is %v", b1, float64(a1/run1), float64(a2/run2), x1, x2)
		}
		return points
	}

	if det == 0 {
		return parallelPoints()
		// parallel lines
	}

	x := (b2*c1 - b1*c2)/det
	y := (a1*c2 - a2*c1)/det

	if isWithin(x, y) {
		return []*Point{{int(x), int(y)}}
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
		if *potential == origin {
			continue
		}
		distance := Manhattan(origin, *potential)
		if distance < currentDistance {
			currentDistance = distance
		}
	}
	return currentDistance
}

func FindCheapest(origin Point, paths []WirePath,candidates ...*Point) int {
	if len(candidates) == 0 || len(paths) != 2 {
		return -1
	}
	currentCost := int(^uint(0) >> 1 )
	fmt.Printf("Calculating cost for %v intersections\n", len(candidates))
	for _, potential := range candidates {
		if *potential == origin {
			continue
		}
		fmt.Println("Calculating cost for intersection: ", potential)
		cost := calculatePathCost(paths[0], potential) + calculatePathCost(paths[1], potential)
		fmt.Println("totalCost: ", cost)
		if cost < currentCost {
			currentCost = cost
		}
	}
	return currentCost
}

func calculatePathCost(path WirePath, intersection *Point) int {
	fmt.Println("calculating path cost")
	pathCost := 0
	previous := Point{0,0}
	for _, step := range path {
		if onPath(previous, step, *intersection) {
			fmt.Println("Ixn on path!")
			cost := partialCost(previous, step, intersection)
			fmt.Println("Returning cost", pathCost, cost)
			pathCost += partialCost(previous, step, intersection)
			break
		}
		fmt.Println("Ixn not on path!")
		cost := partialCost(previous, step, nil)
		fmt.Println("calculated cost: ", cost)
		pathCost += cost
		previous = step
	}
	fmt.Println("Path cost is: ", pathCost)
	return pathCost
}

func onPath(start, end, intersection Point) bool {
	fmt.Printf("Checking if %v is between %v and %v\n", intersection, start, end)
	if start.X == end.X {
		// moving vertically
		min := int(math.Min(float64(start.Y), float64(end.Y)))
		max := int(math.Max(float64(start.Y), float64(end.Y)))
		return intersection.X == start.X &&
			intersection.Y >= min &&
			intersection.Y <= max
	} else {
		// moving horizontally
		min := int(math.Min(float64(start.X), float64(end.X)))
		max := int(math.Max(float64(start.X), float64(end.X)))
		return intersection.Y == start.Y &&
			intersection.X >= min &&
			intersection.X <= max
	}
}

func partialCost(start,end Point, intersection *Point) int {
	if intersection == nil {
		return Manhattan(start, end)
	}
	return Manhattan(start, *intersection)
}
