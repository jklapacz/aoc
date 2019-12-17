package computer_test

import (
	"testing"

	c "github.com/jklapacz/aoc/computer"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert.Equal(t, c.OpcodeAdd, c.Decode(1))
	assert.Equal(t, c.OpcodeAdd, c.Decode(1001))
	assert.Equal(t, c.OpcodeMultiply, c.Decode(1002))
	assert.Equal(t, c.OpcodeMultiply, c.Decode(2))
	assert.Equal(t, c.OpcodeErr, c.Decode(99))
	assert.Equal(t, c.OpcodeErr, c.Decode(1099))
	assert.Equal(t, c.OpcodeUnknown, c.Decode(42))
	assert.Equal(t, c.OpcodeUnknown, c.Decode(9))
}

func TestOps(t *testing.T) {
	basicProgram := "1,0,1,1,99"
	comp := c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assert.Equal(t, "[1 1 1 1 99]", comp.DumpMemory())
	basicProgram = "10101,3,1,0,99"
	comp = c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assert.Equal(t, "[6 3 1 0 99]", comp.DumpMemory())

	basicProgram = "10201,3,1,2,99"
	comp = c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assert.Equal(t, "[10201 3 5 2 99]", comp.DumpMemory())
}
