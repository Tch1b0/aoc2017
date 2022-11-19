package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Tch1b0/polaris/input"
)

func getInput() [][]int {
	return input.Process("./days/02/input.txt", func(str string) [][]int {
		lines := [][]int{}
		str = strings.ReplaceAll(str, "\t", " ")
		str = strings.ReplaceAll(str, "\r", "")
		for _, line := range strings.Split(str, "\n") {
			cols := []int{}
			for _, col := range strings.Split(line, " ") {
				colNum, err := strconv.Atoi(col)
				if err != nil {
					panic(err)
				}
				cols = append(cols, colNum)
			}

			lines = append(lines, cols)
		}

		return lines
	})
}

func min(arr []int) int {
	if len(arr) == 0 {
		panic("Invalid array length")
	}

	var m int = arr[0]

	for _, val := range arr {
		if val < m {
			m = val
		}
	}

	return m
}

func max(arr []int) int {
	if len(arr) == 0 {
		panic("Invalid array length")
	}

	var m int = arr[0]

	for _, val := range arr {
		if val > m {
			m = val
		}
	}

	return m
}

func part1() int {
	sum := 0
	for _, row := range getInput() {
		sum += max(row) - min(row)
	}

	return sum
}

func part2() int {
	return -1
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
