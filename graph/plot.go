package graph

import (
	"fmt"

	"github.com/gookit/color"
)

func Plot(points *PointSet) {
	grid := "===== grid ======\n"
	topLeft, bottomRight := points.boundaries()
	red := color.FgRed.Render
	blue := color.FgBlue.Render
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		grid += "|"
		for x := topLeft.X; x <= bottomRight.X; x++ {
			if (points.Origin == Point{x, y}) {
				grid += fmt.Sprintf("%s ", blue("+"))
			} else if points.Contains(Point{x, y}) {
				grid += fmt.Sprintf("%s ", red("+"))
			} else {
				grid += fmt.Sprintf(". ")
			}
		}
		grid += "|\n"
	}
	trace.Println(grid)
}
