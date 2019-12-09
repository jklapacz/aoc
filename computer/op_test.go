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
