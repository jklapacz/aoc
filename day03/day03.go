package day03

import (
	"bufio"
	"fmt"
	"io"
)

func Solve(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return 1
}
