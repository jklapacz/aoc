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
	assert.Equal(t, c.OpcodeRel, c.Decode(9))
}

func TestOps(t *testing.T) {
	basicProgram := "1,0,1,1,99"
	comp := c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assertMemoryIs(t, comp, 1, 1)

	basicProgram = "10101,3,1,0,99"
	comp = c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assertMemoryIs(t, comp, 0, 6)

	basicProgram = "10201,3,1,2,99"
	comp = c.CreateComputer(basicProgram, nil, nil)
	comp.RunProgram()
	assertMemoryIs(t, comp, 2, 5)

	basicProgram = "203,4,99,0,1"
	comp = c.CreateComputer(basicProgram, nil, nil, 69)
	comp.RunProgram()
	assertMemoryIs(t, comp, 4, 69)

	basicProgram = "209,0,203,1,99,0,1,12"
	comp = c.CreateComputer(basicProgram, nil, nil, 69)
	comp.RunProgram()
	assertMemoryIs(t, comp, 210, 69)
}
