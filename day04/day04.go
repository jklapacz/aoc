package day04

import (
	"fmt"
	"math"
	"strconv"
)

type Passcode struct {
	literal string
	numeric *int
}

type Range struct {
	Start, End int
}

func TotalValidPasscodesInRange(start, end int) int {
	total := 0
	r := Range{start, end}
	for current := start; current <= end; current++ {
		passcodeLiteral := fmt.Sprint(current)
		if SatisfactoryPasscode(passcodeLiteral, r) {
			total++
		}
	}
	return total
}

func (p Passcode) Value() int {
	if p.numeric != nil {
		return *p.numeric
	}
	n, err := strconv.ParseInt(p.literal, 10, 64)
	if err != nil {
		fmt.Println("error converting to num", err)
		return math.MinInt64
	} else {
		intN := int(n)
		p.numeric = &intN
		return *p.numeric
	}
}

func makePasscode(p string) Passcode {
	passcode := Passcode{
		literal: p,
	}
	return passcode
}

func SatisfactoryPasscode(s string, r Range) bool {
	p := makePasscode(s)
	//It is a six-digit number.
	//	The value is within the range given in your puzzle input.
	//	Two adjacent digits are the same (like 22 in 122345).
	//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	if !criteriaLength(p) { return false }
	if !criteriaInRange(p, r) { return false }
	if !criteriaDoubles(p) { return false }
	if !criteriaIncreasing(p) { return false }
	return true
}

func criteriaLength(p Passcode) bool {
	return len(p.literal) == 6
}

func criteriaInRange(p Passcode, r Range) bool {
	return p.Value() >= r.Start && p.Value() <= r.End
}

func criteriaDoubles(p Passcode) bool {
	// check if there is a double character
	valuecounts := make(map[int32]int, len(p.literal))
	for _, c := range p.literal {
		valuecounts[c]++
	}

	previous := int32(p.literal[0])
	for _, current := range p.literal[1:] {
		if current == previous && valuecounts[current] == 2 {
			return true
		}
		previous = current
	}
	return false
}

func criteriaIncreasing(p Passcode) bool {
	currentMin := int32(p.literal[0])
	for _, c := range p.literal {
		if c < currentMin {
			return false
		}
		currentMin = c
	}
	return true
}
