package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const TOO_LONG = 10

func getInput() []uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	var input []uint64
	for scanner.Scan() {
		if scanner.Text() != "" {
			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalln(err)
			}
			input = append(input, uint64(n))
		}
	}
	return input
}

func main() {
	input := getInput()
	fmt.Println(input)
}

