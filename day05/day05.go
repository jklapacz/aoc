package day05

import (
	"fmt"
	"github.com/jklapacz/aoc/day02"
	"log"
	"strconv"
	"strings"
)

type Operation struct {
	opcode day02.Opcode
	encoded int
	output int
	params []int
}

var OpLengths = map[day02.Opcode]int{
	day02.OpcodeAdd: 4,
	day02.OpcodeMultiply: 4,
	day02.OpcodeSave: 2,
	day02.OpcodeOutput: 2,
}

func ParseOperation(opcodes []int, instructionIdx int) Operation{
	encodedOp := opcodes[instructionIdx]
	//fmt.Println("parsing operation at ", instructionIdx, encodedOp)
	op := Operation{opcode: decodeOp(encodedOp), encoded: encodedOp}
	switch op.opcode {
	case day02.OpcodeAdd:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1:instructionIdx+4]
	case day02.OpcodeMultiply:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1:instructionIdx+4]
	case day02.OpcodeOutput:
		op.output = -1
		op.params = opcodes[instructionIdx+1:instructionIdx+2]
	case day02.OpcodeSave:
		op.params= opcodes[instructionIdx+1:instructionIdx+2]
	case day02.OpcodeErr:
		fmt.Println("\t== failing gracefully")
	default:
		log.Fatal("unknown opcode")
	}
	return op
}

func decodeOp(encoded int) day02.Opcode {
	tens := nthdigit(encoded, 1)
	ones:= nthdigit(encoded, 0)
	if tens == 9 && ones == 9 {
		return day02.OpcodeErr
	}
	if ones < 1 || ones > 4 {
		return day02.OpcodeUnknown
	}
	return ones
}

func RunComputer(input string) {
	ops := convertRawInput(input)
	instructionIdx := 0
	for {
		log.Println("\t== ", instructionIdx)
		currentOperation := ParseOperation(ops, instructionIdx)
		if currentOperation.opcode == day02.OpcodeErr {
			log.Fatal("Termination sequence activated")
		}
		performOperation(ops, currentOperation, 1)
		//fmt.Println("== compute loop", currentOperation, instrOffset)
		instructionIdx += OpLengths[currentOperation.opcode]
	}
}

func performOperation(instructions []int, op Operation, input int) {
	immediateA := nthdigit(op.encoded, 2)
	immediateB := nthdigit(op.encoded, 3)

	//fmt.Printf("\t====Immediate settings: a: %v b: %v c: %v\n", immediateA, immediateB, immediateC)

	var paramA, paramB int
	if immediateA == 1 {
		//fmt.Println("!")
		paramA = op.params[0]
		paramA = op.params[0]
	} else {
		//fmt.Println("!!!")
		paramA = instructions[op.params[0]]
	}

	if len(op.params) > 1 {
		if immediateB == 1 {
			paramB = op.params[1]
		} else {
			paramB = instructions[op.params[1]]
		}
	}

	//fmt.Printf("\t====Params: a: %v b: %v c: %v\n", paramA, paramB, paramC)
	//fmt.Printf("\t====op params: %v\n", op.params)
	//fmt.Println("=== running command: ", op, instructions)

	switch op.opcode {
	case day02.OpcodeAdd:
		if len(op.params) < 3 {
			fmt.Println("bad input!")
			return
		}
		instructions[op.params[2]] = paramA + paramB
	case day02.OpcodeMultiply:
		if len(op.params) < 3 {
			fmt.Println("bad input!")
			return
		}
		instructions[op.params[2]] = paramA * paramB
	case day02.OpcodeSave:
		instructions[op.params[0]] = input
	case day02.OpcodeOutput:
		fmt.Println("Output of print command: ", instructions[op.params[0]])
	default:
		log.Fatal("Unsupported opcode!")
	}
}

func nthdigit(x, n int) int {
	powersof10 := []int{1, 10, 100, 1000, 10000}
	return ((x / powersof10[n]) % 10)
}



func convertRawInput(input string) []int {
	ops := strings.Split(input, ",")
	var output []int
	for _, op := range ops {
		opVal, err := strconv.ParseInt(op, 10, 64)
		if err != nil {
			log.Fatal("Op is not an int!")
		}
		output = append(output, int(opVal))
	}
	return output
}
