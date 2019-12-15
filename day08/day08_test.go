package day08_test

import (
	"testing"

	"github.com/jklapacz/aoc/image"
	"github.com/jklapacz/aoc/puzzle"
	"github.com/stretchr/testify/assert"
)

var p *puzzle.Puzzle

func init() {
	p = &puzzle.Puzzle{Filename: "input.txt"}
}

func TestPart01(t *testing.T) {
	cleanup := p.SetupPuzzle()
	defer cleanup()

	assert.Equal(t, 1905, image.Checksum(p.ToString(), 25, 6))
}

func TestPart02(t *testing.T) {
	cleanup := p.SetupPuzzle()
	defer cleanup()
	i := image.CreateImage(p.ToString(), 25, 6)
	i.Decode()
	i.Print()
	i.ToPNG()
	assert.Equal(t, 1, 1)
}
