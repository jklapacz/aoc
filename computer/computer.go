package computer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operation struct {
	opcode          Opcode
	encoded         int
	output          int
	params          []int
	nextInstruction int
}

func ParseOperation(opcodes []int, instructionIdx int) Operation {
	encodedOp := opcodes[instructionIdx]
	//fmt.Println("parsing operation at ", instructionIdx, encodedOp)
	op := Operation{opcode: Decode(encodedOp), encoded: encodedOp}
	switch op.opcode {
	case OpcodeAdd:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeMultiply:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeOutput:
		op.output = -1
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeSave:
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeJIT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeJIF:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeLT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeEq:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	case OpcodeErr:
		fmt.Println("\t== failing gracefully")
	default:
		log.Fatal("unknown opcode", op.opcode)
	}
	return op
}

func RunComputer(input string, inputVal int) int {
	ops := convertRawInput(input)
	instructionIdx := 0
	var output int
	for {
		log.Println("\t== ", instructionIdx)
		currentOperation := ParseOperation(ops, instructionIdx)
		if currentOperation.opcode == OpcodeErr {
			return output
		}
		outputMaybe, next := performOperation(ops, currentOperation, inputVal)
		if outputMaybe != nil {
			output = *outputMaybe
		}
		//fmt.Println("== compute loop", currentOperation, instrOffset)
		fmt.Println("\tnext destination:", currentOperation.nextInstruction)

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
	case OpcodeAdd:
		instructions[op.params[2]] = paramA + paramB
	case OpcodeMultiply:
		instructions[op.params[2]] = paramA * paramB
	case OpcodeSave:
		instructions[op.params[0]] = input
	case OpcodeJIT:
		log.Printf("\t====Jumping to %v if true! %v\n", paramB, paramA)
		if paramA != 0 {
			op.nextInstruction = paramB
		}
	case OpcodeJIF:
		log.Printf("\t====Jumping to %v if false! %v\n", paramB, paramA)
		if paramA == 0 {
			op.nextInstruction = paramB
			fmt.Println("\t\t false!", op)
		}
	case OpcodeLT:
		if paramA < paramB {
			instructions[op.params[2]] = 1
		} else {
			instructions[op.params[2]] = 0
		}
	case OpcodeEq:
		if paramA == paramB {
			instructions[op.params[2]] = 1
		} else {
			instructions[op.params[2]] = 0
		}
	case OpcodeOutput:
		fmt.Println("aaa", paramA)

		fmt.Println("Output of print command: ", paramA) //instructions[op.params[0]])
		return &paramA, op.nextInstruction
		//return &instructions[op.params[0]], op.nextInstruction
	default:
		log.Fatal("Unsupported opcode!")
	}
	return nil, op.nextInstruction
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
