package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(input[0])
	for _, v := range input[1 : n+1] {
		if len(v) > TOO_LONG {
			fmt.Printf("%s%d%s\n", string(v[0]), len(v)-2, string(v[len(v)-1]))
			continue
		}
		fmt.Println(v)
	}
}
