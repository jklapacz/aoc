package computer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IO struct {
	data []int
}

type memoryAddress int

type Computer struct {
	Program         *IO
	UserInput       int
	functionPointer memoryAddress
	output          int
}

func CreateComputer(input string) *Computer {
	programData := &IO{convertRawInput(input)}
	c := &Computer{Program: programData}
	return c
}

// RunProgram runs the computer and returns the output
func (c *Computer) RunProgram() int {
	ops := c.Program.data
	inputVal := c.UserInput
	var output int
	for {
		currentOperation := c.getCurrentOperation()
		if currentOperation.opcode == OpcodeErr {
			return output
		}
		outputMaybe, next := performOperation(ops, currentOperation, inputVal)
		if outputMaybe != nil {
			output = *outputMaybe
		}
		fmt.Println("\tnext destination:", currentOperation.nextInstruction)

		c.functionPointer = memoryAddress(next)
	}
	return output
}

func (c *Computer) GetUserInput(input int) {
	c.UserInput = input
}

func (io *IO) store(value, target int) {
	if target >= 0 && target <= len(io.data) {
		io.data[target] = value
	}
}

func (io *IO) read(target int) int {
	if target >= 0 && target <= len(io.data) {
		return io.data[target]
	}
	log.Fatal("Reading out of bounds data at index: ", target)
	return 0
}

func (c *Computer) getCurrentOperation() Operation {
	return ParseOperation(c.Program.data, c.functionPointer)
}

func ParseOperation(opcodes []int, address memoryAddress) Operation {
	instructionIdx := int(address)
	encodedOp := opcodes[instructionIdx]
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

func performOperation(instructions []int, op Operation, input int) (*int, int) {
	immediateA := nthdigit(op.encoded, 2)
	immediateB := nthdigit(op.encoded, 3)

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
