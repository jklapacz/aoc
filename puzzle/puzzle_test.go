package puzzle_test

import (
	"github.com/jklapacz/aoc/puzzle"
	"testing"
)

const (
	InputFile = "test_data.txt"
)

func TestPuzzle_Setup(t *testing.T) {
	testPuzzle := &puzzle.Puzzle{Filename: InputFile}
	testPuzzle.Setup()
	t.Log("puzzle setup completed without failure")
	testPuzzle.DumpContents()
	t.Log("puzzle contents dumped without failure")
}

func TestPuzzle_SetupBadInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic!")
		}
	}()
	testPuzzle := &puzzle.Puzzle{}
	testPuzzle.Setup()
}
