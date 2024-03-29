package puzzle

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Puzzle struct {
	*os.File
	Filename string
	Instance *Instance
}

type Instance = bufio.Scanner

func (p *Puzzle) Setup() func() {
	return p.SetupPuzzle()
}

func (p *Puzzle) SetupPuzzle() func() {
	if p.Filename == "" {
		panic("no file given!")
	}
	file, err := os.Open(p.Filename)
	if err != nil {
		panic(fmt.Sprintf("could not open file!: %v", err))
	}
	p.Instance = instantiatePuzzle(file)
	p.File = file
	return func() { p.File.Close() }
}

func instantiatePuzzle(file *os.File) *Instance {
	return bufio.NewScanner(file)
}

// mostly for debugging, iterate through the file
func (p *Puzzle) DumpContents() {
	if p.Instance == nil {
		log.Fatal("puzzle is not instantiated!")
	}
	for p.Instance.Scan() {
		fmt.Println(p.Instance.Text())
	}
}

func (p *Puzzle) ToString() string {
	for p.Instance.Scan() {
		return p.Instance.Text()
	}
	return ""
}
