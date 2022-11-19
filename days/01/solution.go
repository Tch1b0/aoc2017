package main

import (
	"fmt"

	"github.com/Tch1b0/polaris/input"
)

func getInput() string {
	return input.Process("days/01/input.txt", func(in string) string {
		newStr := ""
		for _, chr := range in {
			// only accept runes that digits
			if chr >= '0' && chr <= '9' {
				newStr += string(chr)
			}
		}
		return newStr
	})
}

func part1() int {
	in := getInput()
	prev := ' '
	sum := 0
	if in[0] == in[len(in)-1] {
		sum += int(in[0]) - int('0')
	}
	for _, digit := range in {
		if prev != ' ' {
			if prev == digit {
				sum += int(digit) - int('0')
				prev = ' '
			}
			prev = digit
		} else {
			prev = digit
		}
	}

	return sum
}

func part2() int {
	in := getInput()
	ni := in[len(in)/2:] + in[:len(in)/2]
	sum := 0

	for i, digit := range in {
		if digit == rune(ni[i]) {
			sum += int(digit) - int('0')
		}
	}

	return sum
}

func main() {
	fmt.Printf("PART 1: Sum is %d\n", part1())

	fmt.Printf("PART 2: Sum is %d\n", part2())
}
