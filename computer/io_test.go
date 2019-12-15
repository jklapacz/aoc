package computer_test

import (
	"testing"

	"github.com/jklapacz/aoc/computer"
	"github.com/stretchr/testify/assert"
)

func TestUserIO_Read(t *testing.T) {
	readingIO := computer.InitIO(nil, nil)
	writingIO := computer.InitIO(nil, readingIO.Input)
	go func() {
		readVal := readingIO.Read()
		t.Logf("\t I read the following value: %v\n", readVal)
		assert.Equal(t, 1234, readVal)
		readVal = readingIO.Read()
		assert.Equal(t, 1010, readVal)
	}()
	writingIO.Write(1234)
	writingIO.Write(1010)
}
