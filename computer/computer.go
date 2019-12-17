package computer

import (
	"log"
	"strconv"
	"strings"
)

type IO struct {
	data []int
}

type memoryAddress int

type programConfig = []int

type Computer struct {
	Program          *IO
	config           programConfig
	UserInputStreams *UserIO
	functionPointer  *memoryAddress
	output           int
	Interrupt        chan int
	relative         int
}

func CreateComputer(input string, userInput, userOutput chan int, config ...int) *Computer {
	program := loadProgram(input)
	startAddress := memoryAddress(0)
	interrupt := make(chan int, 2)
	return &Computer{
		Program:          program,
		config:           config,
		functionPointer:  &startAddress,
		UserInputStreams: InitIO(userInput, userOutput),
		Interrupt:        interrupt,
		relative:         0,
	}
}

func loadProgram(input string) *IO {
	ops := strings.Split(input, ",")
	var output []int
	for _, op := range ops {
		opVal, err := strconv.ParseInt(op, 10, 64)
		if err != nil {
			log.Fatal("Op is not an int!")
		}
		output = append(output, int(opVal))
	}
	return &IO{output}
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
