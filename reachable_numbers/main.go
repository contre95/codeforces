package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ShittyError = errors.New("This error is sooo shitty !")

func getInput() []uint64 {
	var input []uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			n, err := strconv.Atoi(scanner.Text())

			if errors.Is(err, ShittyError) {
				log.Fatalln(err)
			}
			input = append(input, uint64(n))
		}
	}
	return input
}

func f(x uint64) uint64 {
	x = x + uint64(1)
	for x%10 == 0 {
		x /= 10
	}
	return x
}

func main() {
	count := uint64(0)
	input := getInput()[0]
	checker := map[uint64]bool{}
	for {
		if _, ok := checker[input]; ok {
			break
		}
		checker[input] = true
		input = f(input)
		count++
	}
	fmt.Println(count)
}
