package graph

import (
	"sort"
)

type Grid struct {
	TopLeft, BottomRight Point
}

// -140, -720
type PointSet struct {
	TopLeft, BottomRight, Origin Point
	Members                      map[Point]bool
}

func iterateThroughGrid(topLeft, bottomRight Point, apply func(x, y int)) {
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			apply(x, y)
		}
	}
}

func (g *Grid) Enumerate() *PointSet {
	ps := &PointSet{TopLeft: g.TopLeft, BottomRight: g.BottomRight}
	apply := func(x, y int) {
		ps.Add(Point{X: x, Y: y})
	}
	iterateThroughGrid(g.TopLeft, g.BottomRight, apply)
	return ps
}

func (ps *PointSet) Remove(p Point) {
	delete(ps.Members, p)
}

func (ps *PointSet) Add(p Point) {
	if ps.Members == nil {
		ps.Members = make(map[Point]bool, 0)
	}
	ps.Members[p] = true
}

func (ps *PointSet) Contains(p Point) bool {
	_, ok := ps.Members[p]
	return ok
}

func (ps *PointSet) boundaries() (Point, Point) {
	return ps.TopLeft, ps.BottomRight
}

func (ps *PointSet) Map(reducer func(p Point)) {
	for p := range ps.Members {
		reducer(p)
	}
}

func PointsInLine(origin, target, topLeft, bottomRight Point) *PointSet {
	isInBounds := func(p Point) bool {
		return p.X >= topLeft.X &&
			p.X <= bottomRight.X &&
			p.Y >= topLeft.Y &&
			p.Y <= bottomRight.Y
	}
	m := CalculateSlope(origin, target)
	line := CreateLine(origin, target)
	// move right
	trace.Printf("moving in the following direction: %v, slope: %v\n", slopeDirection(m).ToString(), m)
	points := &PointSet{TopLeft: topLeft, BottomRight: bottomRight, Origin: origin}
	if m.Run == 0 {
		if slopeDirection(m) == dirUp {
			for y := origin.Y; y >= topLeft.Y; y-- {
				points.Add(Point{X: origin.X, Y: y})
			}
		} else {
			for y := origin.Y; y <= bottomRight.Y; y++ {
				points.Add(Point{X: origin.X, Y: y})
			}
		}
		return points
	}

	addPossiblePoints := func(x int) {
		possiblePoint := line.CreatePoint(x)
		// fmt.Println("possible: ", possiblePoint)
		if isInBounds(possiblePoint) {
			// fmt.Println("in bounds!")
			points.Add(possiblePoint)
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

type By func(p1, p2 *Point) bool

type pointSorter struct {
	points []Point
	by     func(p1, p2 *Point) bool
}

func (by By) Sort(points []Point) {
	ps := &pointSorter{
		points: points,
		by:     by,
	}
	sort.Sort(ps)
}

func (s *pointSorter) Swap(i, j int) {
	s.points[i], s.points[j] = s.points[j], s.points[i]
}

func (s *pointSorter) Len() int {
	return len(s.points)
}

func (s *pointSorter) Less(i, j int) bool {
	return s.by(&s.points[i], &s.points[j])
}

func SortByAngle(origin Point, points []Point) {
	//var laserSorted []Point
	//seenByAnother := map[Point]bool{}
	//	pointsToSort := make([]Point, len(points))

	//for len(pointsToSort) > 0 {
	angles := []float64{}
	pointSeen := map[Point]struct{}{}
	angleSet := map[float64][]Point{}
	angle := func(p1, p2 *Point) bool {
		angle1 := CalculateAngle(origin, *p1)
		angle2 := CalculateAngle(origin, *p2)
		trace.Printf("angle1 (%v) = %v angle2 (%v) = %v\n", *p1, angle1, *p2, angle2)
		if _, ok := angleSet[angle1]; !ok {
			angleSet[angle1] = []Point{*p1}
			angles = append(angles, angle1)
			pointSeen[*p1] = struct{}{}
		} else {
			if _, ok := pointSeen[*p1]; !ok {
				angleSet[angle1] = append(angleSet[angle1], *p1)
				pointSeen[*p1] = struct{}{}
			}
		}
		if _, ok := angleSet[angle2]; !ok {
			angleSet[angle2] = []Point{*p2}
			angles = append(angles, angle2)
			pointSeen[*p2] = struct{}{}
		} else {
			if _, ok := pointSeen[*p2]; !ok {
				angleSet[angle2] = append(angleSet[angle2], *p2)
				pointSeen[*p2] = struct{}{}
			}
		}
		if angle1 == angle2 {
			angle1 += float64(Manhattan(origin, *p1))
			angle2 += float64(Manhattan(origin, *p2))
			//			if angle1 < angle2 {
			//				seenByAnother[*p1] = true
			//			} else {
			//				seenByAnother[*p2] = true
			//			}
		}
		return angle1 < angle2
	}
	By(angle).Sort(points)
	distance := func(p1, p2 *Point) bool {
		d1 := float64(Manhattan(origin, *p1))
		d2 := float64(Manhattan(origin, *p2))
		return d1 < d2
	}

	for _, as := range angleSet {
		By(distance).Sort(as)
	}
	sort.Float64s(angles)
	//fmt.Println(angles)
	//fmt.Println(angleSet)
	sortedAsteroids := []Point{}
	for len(sortedAsteroids) != len(points) {
		for _, currentAngle := range angles {
			if len(angleSet[currentAngle]) == 0 {
				continue
			}
			sortedAsteroids = append(sortedAsteroids, angleSet[currentAngle][0])
			angleSet[currentAngle] = angleSet[currentAngle][1:]
		}
	}
	copy(points, sortedAsteroids)
	//temp := make([]Point, len(pointsToSort))
	//copy(temp, pointsToSort)
	//for i, p := range temp {
	//	if _, ok := seenByAnother[p]; !ok {
	//		laserSorted = append(laserSorted, p)
	//		pointsToSort[i] = pointsToSort[len(pointsToSort)-1]
	//		pointsToSort[len(pointsToSort)-1] = Point{X: 0, Y: 0}
	//		pointsToSort = pointsToSort[:len(pointsToSort)-1]
	//		fmt.Println(laserSorted)
	//	}
	//}
	//}
}
