package day10_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/jklapacz/aoc/graph"
	"github.com/stretchr/testify/assert"
)

const (
	test0Input = `.#..#
.....
#####
....#
...##`
	test1Input = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	test2Input = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`
	test3Input = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`
	test4Input = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
)

type SpaceCell struct {
	contents string
}

func IsAsteroid(c rune) bool {
	return c == rune('#')
}

type AsteroidMap struct {
	*graph.PointSet
}

func Parse(input string) *AsteroidMap {
	asteroids := &AsteroidMap{&graph.PointSet{}}
	stringReader := strings.NewReader(input)
	scanner := bufio.NewScanner(stringReader)
	lineCount := -1
	var width int
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line) - 1
		lines = append(lines, line)
		lineCount++
	}
	//fmt.Println(width, lineCount)
	asteroids.TopLeft = graph.Point{X: 0, Y: 0}
	asteroids.BottomRight = graph.Point{X: width, Y: lineCount}
	for y, row := range lines {
		for x, c := range row {
			if IsAsteroid(c) {
				asteroids.Add(graph.Point{X: x, Y: y})
			}
		}
	}
	return asteroids
}

func TestParse(t *testing.T) {
	type scenario struct {
		input          string
		dimensions     []graph.Point
		asteroids      []graph.Point
		totalAsteroids int
	}
	scenarios := []scenario{
		{
			input:          test0Input,
			dimensions:     []graph.Point{{X: 0, Y: 0}, {X: 4, Y: 4}},
			asteroids:      []graph.Point{{X: 1, Y: 0}, {X: 3, Y: 4}},
			totalAsteroids: 10,
		},
		{
			input:          test1Input,
			dimensions:     []graph.Point{{X: 0, Y: 0}, {X: 9, Y: 9}},
			asteroids:      []graph.Point{{X: 1, Y: 4}, {X: 8, Y: 0}},
			totalAsteroids: 40,
		},
	}

	for _, s := range scenarios {
		asteroids := Parse(s.input)
		assert.Equal(t, s.dimensions, []graph.Point{asteroids.TopLeft, asteroids.BottomRight})
		assert.Equal(t, s.totalAsteroids, len(asteroids.Members))
		for _, expectedAsteroid := range s.asteroids {
			assert.Equal(t, true, asteroids.Contains(expectedAsteroid))
		}
	}
}

func ExploreUniverse(asteroids *AsteroidMap) graph.Point {
	var mostSeenPoint graph.Point
	var mostSeenCount int
	for asteroid := range asteroids.Members {
		asteroidsSeen := Explore(asteroid, asteroids)
		if asteroidsSeen > mostSeenCount {
			mostSeenPoint = asteroid
			mostSeenCount = asteroidsSeen
		}
	}
	fmt.Printf("at %v we have %v asteroids seen\n", mostSeenPoint, mostSeenCount)
	return mostSeenPoint
}

func duplicate(original *AsteroidMap) *graph.PointSet {
	points := &graph.PointSet{}
	for p := range original.Members {
		points.Add(p)
	}
	return points
}

func Explore(origin graph.Point, asteroids *AsteroidMap) int {
	unseen := duplicate(asteroids)
	unseen.Remove(origin)
	seen := &graph.PointSet{}

	//fmt.Println("unseen: ", unseen.Members)
	for len(unseen.Members) > 0 {
		for target := range unseen.Members {
			lineToAsteroid := graph.PointsInLine(origin, target, asteroids.TopLeft, asteroids.BottomRight)
			var asteroidsInLine []graph.Point
			//			fmt.Println("Seen so far", len(seen.Members))
			for point := range lineToAsteroid.Members {
				if asteroids.Contains(point) && unseen.Contains(point) {
					unseen.Remove(point)
					asteroidsInLine = append(asteroidsInLine, point)
				}
			}
			closestPoint := graph.FindClosest(origin, asteroidsInLine...)
			seen.Add(closestPoint)
		}
	}
	return len(seen.Members)
}

func TestExploreUniverse(t *testing.T) {
	//input := `.#..##`
	universe := Parse(test0Input)
	assert.Equal(t, graph.Point{X: 3, Y: 4}, ExploreUniverse(universe))
	universe = Parse(test1Input)
	assert.Equal(t, graph.Point{X: 5, Y: 8}, ExploreUniverse(universe))
	universe = Parse(test2Input)
	assert.Equal(t, graph.Point{X: 1, Y: 2}, ExploreUniverse(universe))
	universe = Parse(test3Input)
	assert.Equal(t, graph.Point{X: 6, Y: 3}, ExploreUniverse(universe))
	universe = Parse(test4Input)
	assert.Equal(t, graph.Point{X: 11, Y: 13}, ExploreUniverse(universe))
}

const puzzleInput = `.###.#...#.#.##.#.####..
.#....#####...#.######..
#.#.###.###.#.....#.####
##.###..##..####.#.####.
###########.#######.##.#
##########.#########.##.
.#.##.########.##...###.
###.#.##.#####.#.###.###
##.#####.##..###.#.##.#.
.#.#.#####.####.#..#####
.###.#####.#..#..##.#.##
########.##.#...########
.####..##..#.###.###.#.#
....######.##.#.######.#
###.####.######.#....###
############.#.#.##.####
##...##..####.####.#..##
.###.#########.###..#.##
#.##.#.#...##...#####..#
##.#..###############.##
##.###.#####.##.######..
##.#####.#.#.##..#######
...#######.######...####
#....#.#.#.####.#.#.#.##`

func TestPart1(t *testing.T) {
	universe := Parse(puzzleInput)
	assert.Equal(t, graph.Point{X: 3, Y: 4}, ExploreUniverse(universe))
}

func Extract(asteroids *AsteroidMap, exclude ...graph.Point) []graph.Point {
	amap := duplicate(asteroids)
	for _, p := range exclude {
		amap.Remove(p)
	}
	var points []graph.Point
	for point := range amap.Members {
		points = append(points, point)
	}
	return points
}

func TestPart2(t *testing.T) {
	universe := Parse(test4Input)
	base := ExploreUniverse(universe)
	assert.Equal(t, graph.Point{X: 11, Y: 13}, base)
	asteroids := Extract(universe, base)
	graph.SortByAngle(base, asteroids)
	assert.Equal(t, true, len(asteroids) > 200)
	assert.Equal(t, graph.Point{X: 8, Y: 2}, asteroids[199])
	universe = Parse(puzzleInput)
	base = ExploreUniverse(universe)
	base = graph.Point{X: 3, Y: 4}
	asteroids = Extract(universe, base)
	graph.SortByAngle(base, asteroids)
	assert.Equal(t, true, len(asteroids) > 200)
	assert.Equal(t, graph.Point{X: 7, Y: 6}, asteroids[199])
}

// for every asteroids
//	get every OTHER asteroid #unseen asteroids
//		get the line between a->oA
//		check all points for unseen asteroids
//		if point is not in unseen set, delete it from line
//		get closest point in this line to origin, delete from unseen move to seen
//	once no more unseen asteroids: return count of seen asteroids
