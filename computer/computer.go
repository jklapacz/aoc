package computer

import (
	"errors"
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
	Program          *IO
	UserInput        []int
	UserInputStreams *UserIO
	functionPointer  *memoryAddress
	output           int
	Interrupt        chan int
}

func CreateComputer(input string, userInput, userOutput chan int) *Computer {
	programData := &IO{convertRawInput(input)}
	startAddress := memoryAddress(0)
	interrupt := make(chan int, 2)
	return &Computer{
		Program:          programData,
		UserInput:        []int{},
		functionPointer:  &startAddress,
		UserInputStreams: InitIO(userInput, userOutput),
		Interrupt:        interrupt,
	}
}

// RunProgram runs the computer and returns the output
func (c *Computer) RunProgram() {
	for {
		currentOperation := c.getCurrentOperation()
		if currentOperation.opcode == OpcodeErr {
			c.Interrupt <- c.output
			return
		}
		c.performOperation(currentOperation)
	}
}

func ReadFromComputer(c *Computer) (int, error) {
	select {
	case computerOutput := <-c.UserInputStreams.Output:
		return computerOutput, nil
	case output := <-c.Interrupt:
		fmt.Println("beefcake!")
		return output, errors.New("!!")
	}
}

func WriteToComputer(c *Computer, userInput int) {
	c.UserInputStreams.Input <- userInput
}

func (c *Computer) GetUserInput(input int) {
	//log.Printf("== getting user input: %v, past: %v\n", input, c.UserInput)
	c.UserInput = append(c.UserInput, input)
}

func (c *Computer) getInput() int {
	//log.Printf("reading stored input")
	var currentInput int
	if len(c.UserInput) == 0 {
		return c.UserInputStreams.Read()
	} else if len(c.UserInput) == 1 {
		currentInput = c.UserInput[0]
		c.UserInput = []int{}
	} else {
		currentInput = c.UserInput[0]
		c.UserInput = c.UserInput[1:]
	}
	return currentInput
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
	return ParseOperation(c.Program.data, *c.functionPointer)
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
		//fmt.Printf("\t== at address %v", instructionIdx)
		fmt.Println("\t== failing gracefully")
	default:
		log.Fatal("unknown opcode", op.opcode)
	}
	return op
}

func (c *Computer) performOperation(op Operation) {
	immediateA := nthdigit(op.encoded, 2)
	immediateB := nthdigit(op.encoded, 3)

	var paramA, paramB int
	if immediateA == 1 {
		paramA = op.params[0]
	} else {
		paramA = c.Program.read(op.params[0])
	}

	if len(op.params) > 1 {
		if immediateB == 1 {
			paramB = op.params[1]
		} else {
			paramB = c.Program.read(op.params[1])
		}
	}

	switch op.opcode {
	case OpcodeAdd:
		combined := paramA + paramB
		c.Program.store(combined, op.params[2])
	case OpcodeMultiply:
		combined := paramA * paramB
		c.Program.store(combined, op.params[2])
	case OpcodeSave:
		c.Program.store(c.getInput(), op.params[0])
	case OpcodeJIT:
		if paramA != 0 {
			*c.functionPointer = memoryAddress(paramB)
			return
		}
	case OpcodeJIF:
		if paramA == 0 {
			*c.functionPointer = memoryAddress(paramB)
			return
		}
	case OpcodeLT:
		if paramA < paramB {
			c.Program.store(1, op.params[2])
		} else {
			c.Program.store(0, op.params[2])
		}
	case OpcodeEq:
		if paramA == paramB {
			c.Program.store(1, op.params[2])
		} else {
			c.Program.store(0, op.params[2])
		}
	case OpcodeOutput:
		//fmt.Printf("\t=== at instruction: %v", *c.functionPointer)
		//fmt.Printf("\t=== next instruction: %v", op.nextInstruction)
		//fmt.Println("\t Output of print command: ", paramA)
		c.output = paramA
		c.UserInputStreams.Write(paramA)
	default:
		log.Fatal("Unsupported opcode!")
	}
	*c.functionPointer = memoryAddress(op.nextInstruction)
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
