package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Tokenize(filepath string) [][]string {
	file, err := os.Open(filepath)
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
	var tokens [][]string
	for scanner.Scan() {
		line := scanner.Text()
		tokens = append(tokens, strings.Fields(line))
	}
	return tokens
}
