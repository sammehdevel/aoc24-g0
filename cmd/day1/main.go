package main

import (
	"aoc24-go/internal"
	"fmt"
	"math"
	"os"
	"slices"
)

type Day1Data struct {
	left  []int
	right []int
}

func main() {

	var args = os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	var left, right []int
	tokens := internal.Tokenize(args[0])
	for _, line := range tokens {
		var l = internal.ParseInt(line[0])
		var r = internal.ParseInt(line[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)
	dat := Day1Data{left: left, right: right}
	fmt.Printf("Part 1: %d\n", part1(dat))
	fmt.Printf("Part 2: %d\n", part2(dat))
}

func part1(dat Day1Data) int {
	var d = 0
	for i := 0; i < len(dat.left); i++ {
		d += int(math.Abs(float64(dat.left[i] - dat.right[i])))
	}
	return d
}

func part2(dat Day1Data) int {
	var d int
	for i := 0; i < len(dat.left); i++ {
		n := dat.left[i]
		c := 0
		for j := 0; j < len(dat.right); j++ {
			if dat.right[j] == n {
				c++
			}
		}
		d += n * c
	}
	return d
}
