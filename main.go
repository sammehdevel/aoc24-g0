package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var args = os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(1)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		l, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(1)
		}
		left = append(left, l)
		r, err := strconv.Atoi(tokens[len(tokens)-1])
		if err != nil {
			panic(1)
		}
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	var d int
	for i := 0; i < len(left); i++ {
		n := left[i]
		c := 0
		for j := 0; j < len(right); j++ {
			if right[j] == n {
				c++
			}
		}
		d += n * c
	}

	fmt.Println(d)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
