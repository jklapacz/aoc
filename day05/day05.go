package day05

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jklapacz/aoc/day02"
)

type Operation struct {
	opcode          day02.Opcode
	encoded         int
	output          int
	params          []int
	nextInstruction int
}

var OpLengths = map[day02.Opcode]int{
	day02.OpcodeAdd:      4,
	day02.OpcodeMultiply: 4,
	day02.OpcodeSave:     2,
	day02.OpcodeOutput:   2,
	day02.OpcodeJIT:      3,
	day02.OpcodeJIF:      3,
	day02.OpcodeLT:       4,
	day02.OpcodeEq:       4,
}

func ParseOperation(opcodes []int, instructionIdx int) Operation {
	encodedOp := opcodes[instructionIdx]
	//fmt.Println("parsing operation at ", instructionIdx, encodedOp)
	op := Operation{opcode: decodeOp(encodedOp), encoded: encodedOp}
	switch op.opcode {
	case day02.OpcodeAdd:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeMultiply:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeOutput:
		op.output = -1
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeSave:
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeJIT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeJIF:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeLT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeEq:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case day02.OpcodeErr:
		fmt.Println("\t== failing gracefully")
	default:
		log.Fatal("unknown opcode", op.opcode)
	}
	return op
}

func decodeOp(encoded int) day02.Opcode {
	tens := nthdigit(encoded, 1)
	ones := nthdigit(encoded, 0)
	if tens == 9 && ones == 9 {
		return day02.OpcodeErr
	}
	if ones < 1 || ones > 8 {
		return day02.OpcodeUnknown
	}
	return ones
}

func RunComputer(input string, inputVal int) int {
	ops := convertRawInput(input)
	instructionIdx := 0
	var output int
	for {
		log.Println("\t== ", instructionIdx)
		currentOperation := ParseOperation(ops, instructionIdx)
		if currentOperation.opcode == day02.OpcodeErr {
			return output
		}
		outputMaybe, next := performOperation(ops, currentOperation, inputVal)
		if outputMaybe != nil {
			output = *outputMaybe
		}
		//fmt.Println("== compute loop", currentOperation, instrOffset)
		fmt.Println("\tnext destination:", currentOperation.nextInstruction)
		fmt.Println("\tnext destination:", next)

		instructionIdx = next
	}
	return output
}

func performOperation(instructions []int, op Operation, input int) (*int, int) {
	immediateA := nthdigit(op.encoded, 2)
	immediateB := nthdigit(op.encoded, 3)

	fmt.Printf("\t====Immediate settings: a: %v b: %v \n", immediateA, immediateB)

	var paramA, paramB int
	if immediateA == 1 {
		fmt.Println("!")
		paramA = op.params[0]
	} else {
		fmt.Println("!!!")
		paramA = instructions[op.params[0]]
	}

	if len(op.params) > 1 {
		if immediateB == 1 {
			paramB = op.params[1]
		} else {
			paramB = instructions[op.params[1]]
		}
	}

	//fmt.Printf("\t====Params: a: %v b: %v c: %v\n", paramA, paramB)
	//fmt.Printf("\t====op params: %v\n", op.params)
	//fmt.Println("=== running command: ", op, instructions)

	switch op.opcode {
	case day02.OpcodeAdd:
		log.Printf("!! writing to: %v\n", op.params[2])
		instructions[op.params[2]] = paramA + paramB
	case day02.OpcodeMultiply:
		log.Printf("!! writing to: %v\n", op.params[2])
		instructions[op.params[2]] = paramA * paramB
	case day02.OpcodeSave:
		instructions[op.params[0]] = input
	case day02.OpcodeJIT:
		log.Printf("\t====Jumping to %v if true! %v\n", paramB, paramA)
		if paramA != 0 {
			op.nextInstruction = paramB
		}
	case day02.OpcodeJIF:
		log.Printf("\t====Jumping to %v if false! %v\n", paramB, paramA)
		if paramA == 0 {
			op.nextInstruction = paramB
			fmt.Println("\t\t false!", op)
		}
	case day02.OpcodeLT:
		if paramA < paramB {
			instructions[op.params[2]] = 1
		} else {
			instructions[op.params[2]] = 0
		}
	case day02.OpcodeEq:
		if paramA == paramB {
			instructions[op.params[2]] = 1
		} else {
			instructions[op.params[2]] = 0
		}
	case day02.OpcodeOutput:
		fmt.Println("aaa", paramA)

		fmt.Println("Output of print command: ", paramA) //instructions[op.params[0]])
		return &paramA, op.nextInstruction
		//return &instructions[op.params[0]], op.nextInstruction
	default:
		log.Fatal("Unsupported opcode!")
	}
	return nil, op.nextInstruction
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
