package computer_test

import (
	"fmt"
	"testing"

	"github.com/jklapacz/aoc/computer"
	"github.com/stretchr/testify/assert"
)

func assertMemoryIs(t *testing.T, c *computer.Computer, register, expected int) {
	value, err := c.ReadFromMemory(register)
	assert.NoError(t, err)
	assert.Equal(t, expected, value)
}

func TestMemoryAccess(t *testing.T) {
	c := computer.CreateComputer("1234,0,0", nil, nil)
	value, err := c.ReadFromMemory(0)
	assert.NoError(t, err)
	assert.Equal(t, 1234, value)
	assert.Equal(t, 1, 1)

	err = c.WriteToMemory(1, 999)
	assert.NoError(t, err)
	value, err = c.ReadFromMemory(1)
	assert.Equal(t, 999, value)

	assert.Error(t, c.WriteToMemory(4500, 999))

	value, err = c.ReadFromMemory(-1)
	assert.Error(t, err)

	value, err = c.ReadFromMemory(3)
	assert.Error(t, err)
	assert.Contains(t, fmt.Sprintf("%s", err), "empty")

	assert.NoError(t, c.WriteToMemory(3, 421))
	value, err = c.ReadFromMemory(3)
	assert.NoError(t, err)
	assert.Equal(t, 421, value)

	assert.NoError(t, c.WriteToMemory(3, 422))
	value, err = c.ReadFromMemory(3)
	assert.NoError(t, err)
	assert.Equal(t, 422, value)

	//t.Logf(c.Memory.Dump())
}

func TestMemoryOperations(t *testing.T) {
	// Add operation
	program := "21001,0,9,2,99"
	c := computer.CreateComputer(program, nil, nil)
	c.RunProgram()
	assertMemoryIs(t, c, 2, 21010)

	// Multiply operation
	program = "21002,0,9,2,99"
	c = computer.CreateComputer(program, nil, nil)
	c.RunProgram()
	assertMemoryIs(t, c, 2, 189018)

	// Save operation
	program = "203,4,99,0,4"
	c = computer.CreateComputer(program, nil, nil, 987)
	c.RunProgram()
	assertMemoryIs(t, c, 4, 987)
	//t.Logf(c.Memory.Dump())

	// Less than operation
	program = "21007,1,9,2,99"
	c = computer.CreateComputer(program, nil, nil)
	c.RunProgram()
	assertMemoryIs(t, c, 2, 1)

	// EQ operation
	program = "21008,1,1,0,99"
	c = computer.CreateComputer(program, nil, nil)
	c.RunProgram()
	assertMemoryIs(t, c, 0, 1)
}
