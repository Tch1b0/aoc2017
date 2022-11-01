package main

import (
	"fmt"

	"github.com/tch1b0/polaris/input"
)

func getInput() string {
	return input.Process("day_1/input.txt", func(in string) string {
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

func part1() {
	prev := ' '
	sum := 0
	for _, digit := range getInput() {
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

	fmt.Printf("PART 1: Sum is %d\n", sum)
}

func part2() {

}

func main() {
	part1()

	part2()
}
