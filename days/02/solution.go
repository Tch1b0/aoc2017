package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
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

func part1() int {
	sum := 0
	for _, row := range getInput() {
		sum += math.Max(row) - math.Min(row)
	}

	return sum
}

func part2() int {
	sum := 0
	for _, row := range getInput() {
		var s *int = nil
		for _, col := range row {
			for _, divisor := range row {
				if col != divisor && col%divisor == 0 {
					x := col / divisor
					s = &x
					break
				}
			}
			if s != nil {
				break
			}
		}
		sum += *s
	}

	return sum
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}
