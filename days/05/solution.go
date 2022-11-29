package main

import (
    "fmt"
    stdstrings "strings"

    "github.com/Tch1b0/polaris/input"
    "github.com/Tch1b0/polaris/array"
    "github.com/Tch1b0/polaris/strings"
)

func getInput() []int {
    return input.Process("days/05/input.txt", func(str string) []int {
        strs := stdstrings.Split(str, "\n")
        return array.Map(strs[:len(strs)-1], func(v string, i int) int {
            num, err := strings.Atoi(v)
            if err != nil {
                fmt.Println(i)
                panic(err)
            }
            return num
        })
    })
}

func part1() int {
    i := 0
    j := 0
    in := getInput()
    for i >= 0 && i < len(in) {
        v := in[i]
        in[i] += 1
        i = i + v
        j++
    }
    
    return j
}

func part2() int {
    i := 0
    j := 0
    in := getInput()
    for i >= 0 && i < len(in) {
        v := in[i]
        if v >= 3 {
            in[i] -= 1
        } else {
            in[i] += 1
        }
        i = i + v
        j++
    }

    return j
}

func main() {
    fmt.Printf("PART 1: %d\n", part1())

    
    fmt.Printf("PART 2: %d\n", part2())
}

