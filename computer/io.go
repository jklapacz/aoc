package computer

type UserIO struct {
	Input, Output chan int
}

func InitIO(inputIo, outputIo chan int) *UserIO {
	var input, output chan int
	if inputIo != nil {
		input = inputIo
	} else {
		input = make(chan int)
	}
	if outputIo != nil {
		output = outputIo
	} else {
		output = make(chan int)
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
	io.Output <- input
}
