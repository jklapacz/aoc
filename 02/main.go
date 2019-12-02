package main

import (
	"bytes"
	"fmt"
	"github.com/jklapacz/aoc"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	OpcodeAdd = 1
	OpcodeMultiply = 2
	OpcodeErr = 99
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("argument needed")
	}
	filename := os.Args[1]
	puzzle := &aoc.Puzzle{Filename:filename}
	cleanup := puzzle.SetupPuzzle()
	defer cleanup()
	solve(puzzle)
}

type command = *bytes.Buffer


func solve(p *aoc.Puzzle) {
	for p.Instance.Scan() {
		cmd := bytes.NewBufferString(p.Instance.Text())
		execute(cmd)
	}
	fmt.Println("Solving!")
}

func execute(cmd command) {
	rawOperations := strings.Split(cmd.String(), ",")
	operations := make([]int, len(rawOperations))
	for idx, op := range rawOperations {
		parsedVal, err := strconv.ParseInt(op, 10, 32)
		if err != nil {
			log.Fatal("Could not parse opcode")
		}
		operations[idx] = int(parsedVal)
	}
	startingIdx := 0
	offsetIncrement := 4
	for curIdx := startingIdx; curIdx + offsetIncrement < len(operations); curIdx += offsetIncrement {
		fmt.Printf("Next iteration, opcode index: %v \n", curIdx)
		performOperation(operations, curIdx)
	}
	fmt.Println("Output", operations)
}

func performOperation(opcodes []int, opcodeBit int) {
	switch opcodes[opcodeBit] {
	case OpcodeAdd:
		performAddOp(opcodes, opcodeBit)
	case OpcodeMultiply:
		performMultOp(opcodes, opcodeBit)
	case OpcodeErr:
		fmt.Println("Error opcode found")
	default:
		log.Fatal("Unexpected error")
	}
}

func performAddOp(opcodes []int, operationOffset int) {
	IdxAddFirst := opcodes[operationOffset + 1]
	IdxAddSecond := opcodes[operationOffset + 2]
	IdxTarget := opcodes[operationOffset + 3]
	fmt.Printf("Operating on indicies %v, %v and %v\n", IdxAddFirst, IdxAddSecond, IdxTarget)
	fmt.Printf("Value at %v before: %v \n", IdxTarget, opcodes[IdxTarget])
	opcodes[IdxTarget] = opcodes[IdxAddFirst] + opcodes[IdxAddSecond]
	fmt.Printf("Value at %v after: %v \n",IdxTarget, opcodes[IdxTarget])
}

func performMultOp(opcodes []int, operationOffset int) {
	IdxFirst := opcodes[operationOffset + 1]
	IdxSecond := opcodes[operationOffset + 2]
	IdxTarget := opcodes[operationOffset + 3]
	fmt.Printf("Operating on indicies %v, %v and %v\n", IdxFirst, IdxSecond, IdxTarget)
	fmt.Printf("Value at %v before: %v \n", IdxTarget, opcodes[IdxTarget])
	opcodes[IdxTarget] = opcodes[IdxFirst] * opcodes[IdxSecond]
	fmt.Printf("Value at %v after: %v \n",IdxTarget, opcodes[IdxTarget])
}

