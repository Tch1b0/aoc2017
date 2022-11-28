package main

import (
	"fmt"
	stdstrings "strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/strings"
)

func getInput() [][]string {
	return input.Process("days/04/input.txt", func(str string) [][]string {
		var serialized [][]string
		for _, line := range stdstrings.Split(str, "\n") {
			serialized = append(serialized, stdstrings.Split(line, " "))
		}
		return serialized
	})
}

func getValidPasswords(validatePass func(a, b string) bool) int {
	validPasswords := 0
	for _, line := range getInput() {
		isValid := true
		for i, a := range line {
			for j, b := range line {
				if i == j {
					continue
				}
				if validatePass(a, b) {
					isValid = false
				}
			}
		}
		if isValid {
			validPasswords += 1
		}
	}
	return validPasswords
}

func part1() int {
	return getValidPasswords(func(a, b string) bool { return a == b })
}

func part2() int {
	return getValidPasswords(func(a, b string) bool { return a == b || strings.IsAnagram(a, b) })
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
