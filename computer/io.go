package computer

import "fmt"

type UserIO struct {
	Input, Output chan int
}

func InitIO(inputIo, outputIo chan int) *UserIO {
	var input, output chan int
	if inputIo != nil {
		input = inputIo
	} else {
		input = make(chan int, 2)
	}
	if outputIo != nil {
		output = outputIo
	} else {
		output = make(chan int, 2)
	}

	return &UserIO{
		Input:  input,
		Output: output,
	}
}

func (io *UserIO) Read() int {
	select {
	case userInput := <-io.Input:
		return userInput
	}
}

func (io *UserIO) Write(input int) {
	fmt.Printf("=== going to send %v to output\n", input)
	io.Output <- input
}
