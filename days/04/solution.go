package main

import (
	"fmt"
	"strings"

	"github.com/Tch1b0/polaris/input"
)

func getInput() [][]string {
	return input.Process("days/04/input.txt", func(str string) [][]string {
		var serialized [][]string
		for _, line := range strings.Split(str, "\n") {
			serialized = append(serialized, strings.Split(line, " "))
		}
		return serialized
	})
}

func part1() int {
	valid := 0

	for _, line := range getInput() {
		isValid := true
		for i, a := range line {
			for j, b := range line {
				if i == j {
					continue
				}
				if a == b {
					isValid = false
				}
			}
		}
		if isValid {
			valid += 1
		}
	}
	return valid
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
