package day02

import (
	"bytes"
	"fmt"
	aoc "github.com/jklapacz/aoc/puzzle"
	"log"
	"os"
	"strconv"
	"strings"
)

type Opcode = int
const (
	OpcodeAdd Opcode = iota + 1
	OpcodeMultiply
	OpcodeSave
	OpcodeOutput
	OpcodeUnknown Opcode = 98
	OpcodeErr Opcode = 99
)

func SolveDay02() {
	if len(os.Args) < 3 {
		log.Fatal("argument needed")
	}
	filename := os.Args[1]
	part := os.Args[2]
	if part == "2" {
		partTwo(filename)
		return
	}
	puzzle := &aoc.Puzzle{Filename:filename}
	cleanup := puzzle.SetupPuzzle()
	defer cleanup()
	solve(puzzle, 12, 2)
}

func partTwo(filename string) {
	desiredOutput := 19690720
	startNoun := 0
	startVerb := 0
	for noun := startNoun; noun <= 99; noun++ {
		fmt.Println("Searching...")
		for verb := startVerb; verb <= 99; verb++ {
			output := computeOutput(filename, noun, verb)
			if output == desiredOutput {
				fmt.Println("==== desired output found!", noun, verb)
				return
			}
			fmt.Printf("\t[%v] - [%v] inconclusive\n", noun, verb)
		}
	}
	fmt.Println("search space exhausted...")
}

func computeOutput(filename string, noun, verb int) int {
	puzzle := &aoc.Puzzle{Filename:filename}
	cleanup := puzzle.SetupPuzzle()
	defer cleanup()
	return solve(puzzle, noun, verb)
}

type command = *bytes.Buffer


func solve(p *aoc.Puzzle, noun, verb int) int {
	for p.Instance.Scan() {
		cmd := bytes.NewBufferString(p.Instance.Text())
		return ExecuteCommand(cmd, noun, verb)
	}
	return -1
}

func ExecuteCommand(cmd command, noun, verb int) int {
	rawOperations := strings.Split(cmd.String(), ",")
	operations := make([]int, len(rawOperations))
	for idx, op := range rawOperations {
		parsedVal, err := strconv.ParseInt(op, 10, 32)
		if err != nil {
			log.Fatal("Could not parse opcode")
		}
		operations[idx] = int(parsedVal)
	}
	if noun >= 0 {
		operations[1] = noun
	}
	if verb >= 0 {
		operations[2] = verb
	}
	startingIdx := 0
	offsetIncrement := 4
	for curIdx := startingIdx; curIdx + offsetIncrement < len(operations); curIdx += offsetIncrement {
		fmt.Printf("Next iteration, opcode index: %v \n", curIdx)
		performOperation(operations, curIdx)
	}
	fmt.Println("Output", operations)
	return operations[0]
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
		log.Fatal("unknown opcode")
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

