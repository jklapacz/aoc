package computer

import (
	"errors"
	"fmt"
	"log"
)

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

func ContinuouslyRead(c *Computer) []int {
	var outputs []int
	for {
		output, err := ReadFromComputer(c)
		if err != nil {
			return outputs
		}
		outputs = append(outputs, output)
	}
}

func ReadFromComputer(c *Computer) (int, error) {
	select {
	case computerOutput := <-c.UserInputStreams.Output:
		return computerOutput, nil
	case output, _ := <-c.Interrupt:
		fmt.Println("beefcake!")
		return output, errors.New("!!")
	}
}

func WriteToComputer(c *Computer, userInput int) {
	c.UserInputStreams.Input <- userInput
}

func (c *Computer) getInput() int {
	var currentInput int
	if len(c.config) == 0 {
		return c.UserInputStreams.Read()
	} else if len(c.config) == 1 {
		currentInput = c.config[0]
		c.config = []int{}
	} else {
		currentInput = c.config[0]
		c.config = c.config[1:]
	}
	return currentInput
}

func (io *IO) store(value, target int) {
	log.Printf("\t-- storing %v at idx %v\n", value, target)
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
