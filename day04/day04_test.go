package day04_test

import (
	"github.com/jklapacz/aoc/day04"
	"testing"
)

func TestSatisfactoryPasscode(t *testing.T) {
	type scenarios struct {
		input string
		output bool
	}
	testCases := []scenarios{
		{"111122", true },
		{"12345", false},
		{ "223450", false },
		{ "123789", false },
	}

	for _, s := range testCases {
		isSatisfactory := day04.SatisfactoryPasscode(s.input, day04.Range{0, 10000000})
		if isSatisfactory != s.output {
			t.Logf("case: %v\n", s.input)
			t.Logf("passcode satisfaction not met, expected: %v, actual: %v\n", s.output, isSatisfactory)
			t.Fail()
		}
	}
}

func TestTotalValidPasscodesInRange(t *testing.T) {
	type scenarios struct {
		start, end, output int
	}
	testCases := []scenarios{
		{111122, 111133, 2 },
		{12345, 12345, 0 },
		{222222, 222233, 1 },
	}

	for _, s := range testCases {
		total := day04.TotalValidPasscodesInRange(s.start, s.end)
		if total != s.output {
			t.Logf("case: %v-%v\n", s.start, s.end)
			t.Logf("total not found, expected: %v, actual: %v\n", s.output, total)
			t.Fail()
		}
	}
}

//111111 meets these criteria (double 11, never decreases).
//223450 does not meet these criteria (decreasing pair of digits 50).
//123789 does not meet these criteria (no double).
