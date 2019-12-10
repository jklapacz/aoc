package day07_test

import (
	"testing"

	"github.com/jklapacz/aoc/computer"
	"github.com/stretchr/testify/assert"
)

type amplifier struct {
	*computer.Computer
}

func runScenario(input string, phases []int) int {
	amplifierA := createAmplifier(input, phases[0], 0)
	outputA := amplifierA.RunProgram()
	amplifierB := createAmplifier(input, phases[1], outputA)
	outputB := amplifierB.RunProgram()
	amplifierC := createAmplifier(input, phases[2], outputB)
	outputC := amplifierC.RunProgram()
	amplifierD := createAmplifier(input, phases[3], outputC)
	outputD := amplifierD.RunProgram()
	amplifierE := createAmplifier(input, phases[4], outputD)
	return amplifierE.RunProgram()
}

func createAmplifier(input string, phaseSetting, instruction int) *amplifier {
	ampComputer := computer.CreateComputer(input)
	ampComputer.GetUserInput(phaseSetting)
	ampComputer.GetUserInput(instruction)
	return &amplifier{ampComputer}
}

func TestDay07(t *testing.T) {
	assert.Equal(t, 43210, runScenario("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", []int{4, 3, 2, 1, 0}))
	assert.Equal(t, 54321, runScenario("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", []int{0, 1, 2, 3, 4}))
	assert.Equal(t, 65210, runScenario("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", []int{1, 0, 4, 3, 2}))
}
