package day07_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/jklapacz/aoc/computer"
	"github.com/stretchr/testify/assert"
)

func init() {
	runtime.GOMAXPROCS(1)
}

type amplifier struct {
	computer       *computer.Computer
	next           *amplifier
	inFeedbackLoop bool
}

func runScenario(input string, phases []int, inputSetting int) int {
	//outputA := computer.ReadFromComputer(amplifierA.computer)

	//amplifierB := createAmplifier(input, phases[1], outputA)
	//outputB := computer.ReadFromComputer(amplifierB.computer)

	//amplifierC := createAmplifier(input, phases[2], outputB)
	//outputC := computer.ReadFromComputer(amplifierC.computer)

	//amplifierD := createAmplifier(input, phases[3], outputC)

	var ampAInput int
	amplifierA := createChainedAmplifier(input, phases[0], nil, nil)
	amplifierB := createChainedAmplifier(input, phases[1], amplifierA, nil)
	amplifierC := createChainedAmplifier(input, phases[2], amplifierB, nil)
	amplifierD := createChainedAmplifier(input, phases[3], amplifierC, nil)
	amplifierE := createChainedAmplifier(input, phases[4], amplifierD, nil)
	go amplifierA.computer.RunProgram()
	go amplifierB.computer.RunProgram()
	go amplifierC.computer.RunProgram()
	go amplifierD.computer.RunProgram()
	go amplifierE.computer.RunProgram()
	for {
		// sends initial "GO" signal
		computer.WriteToComputer(amplifierA.computer, ampAInput)
		//outputD := computer.ReadFromComputer(amplifierD.computer)

		//amplifierE := createAmplifier(input, phases[4], outputD)

		outputE, err := computer.ReadFromComputer(amplifierE.computer)
		if amplifierA.inFeedbackLoop {
			if err != nil {
				fmt.Println("Err not nil!, returning", outputE)
				return outputE
			} else {
				ampAInput = outputE
			}
		} else {
			return outputE
		}
	}
}

func createChainedAmplifier(input string, phaseSetting int, inputAmp, outputAmp *amplifier) *amplifier {
	var newAmpInputChan, newAmpOutputChan chan int
	if inputAmp != nil {
		newAmpInputChan = inputAmp.computer.UserInputStreams.Output
	}

	if outputAmp != nil {
		newAmpOutputChan = outputAmp.computer.UserInputStreams.Input
	}

	ampComputer := computer.CreateComputer(input, newAmpInputChan, newAmpOutputChan)
	computer.WriteToComputer(ampComputer, phaseSetting)
	feedbackLoopMode := phaseSetting > 4 && phaseSetting <= 9
	return &amplifier{computer: ampComputer, next: nil, inFeedbackLoop: feedbackLoopMode}
}

func createAmplifier(input string, phaseSetting, instruction int) *amplifier {
	ampComputer := computer.CreateComputer(input, nil, nil)
	ampComputer.GetUserInput(phaseSetting)
	ampComputer.GetUserInput(instruction)
	feedbackLoopMode := phaseSetting > 4 && phaseSetting <= 9
	return &amplifier{computer: ampComputer, next: nil, inFeedbackLoop: feedbackLoopMode}
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func TestIterateCombinations(t *testing.T) {
	assert.Equal(t, 120, len(permutations([]int{0, 1, 2, 3, 4})))
}

func TestDay07(t *testing.T) {
	assert.Equal(t, 43210, runScenario("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", []int{4, 3, 2, 1, 0}, 0))
	assert.Equal(t, 54321, runScenario("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", []int{0, 1, 2, 3, 4}, 0))
	assert.Equal(t, 65210, runScenario("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", []int{1, 0, 4, 3, 2}, 0))
}

func getMaxiumum(input string, phases []int) int {
	var maxVal int
	for _, permutation := range permutations(phases) {
		currentVal := runScenario(input, permutation, 0)
		if currentVal > maxVal {
			maxVal = currentVal
		}
	}
	return maxVal
}

func TestDay07Actual(t *testing.T) {
	input := `3,8,1001,8,10,8,105,1,0,0,21,34,47,72,93,110,191,272,353,434,99999,3,9,102,3,9,9,1001,9,3,9,4,9,99,3,9,102,4,9,9,1001,9,4,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,1001,9,2,9,1002,9,2,9,101,4,9,9,4,9,99,3,9,1002,9,3,9,101,5,9,9,102,4,9,9,1001,9,4,9,4,9,99,3,9,101,3,9,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99`
	assert.Equal(t, 17440, getMaxiumum(input, []int{0, 1, 2, 3, 4}))
	assert.Equal(t, 27561242, getMaxiumum(input, []int{5, 6, 7, 8, 9}))
}

func TestDay07Part2(t *testing.T) {
	assert.Equal(t, 139629729, runScenario("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", []int{9, 8, 7, 6, 5}, 0))
	assert.Equal(t, 18216, runScenario("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", []int{9, 7, 8, 5, 6}, 0))
}
