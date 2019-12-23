package graph

import "sort"

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
	angle := func(p1, p2 *Point) bool {
		angle1 := CalculateAngle(origin, *p1)
		angle2 := CalculateAngle(origin, *p2)
		if angle1 == angle2 {
			distance1 := Manhattan(origin, *p1)
			distance2 := Manhattan(origin, *p2)
			return distance1 < distance2
		}
		return angle1 < angle2
	}
	By(angle).Sort(points)
}
