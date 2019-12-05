package day02_test

import (
	"bytes"
	"github.com/jklapacz/aoc/day02"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	type scenario struct {
		input string
		output int
	}
	scenarios := []scenario{
		{ "3,0,4,0,99", 50 },
	}
	for _, s := range scenarios {
		cmdBuffer := bytes.NewBufferString(s.input)
		actual := day02.ExecuteCommand(cmdBuffer, -1, -1)
		if actual != s.output {
			t.Logf("test case %v failed, expected: %v actual %v\n", s.input, s.output, actual)
			t.Fail()
		}
	}
}
