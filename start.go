package main

import (
	"bufio"
	"fmt"
	"os"
)

const TOO_LONG = 10

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
func main() {
	input := getInput()
	fmt.Println(input)
}
