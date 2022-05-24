package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

//func countZeros(i uint64) int {
//count := 0
//for i%10 == 0 {
//i = i / 10
//count++
//}
//return count
//}

//func countDigits(i uint64) uint64 {
//count := uint64(0)
//for i >= 1 {
//i = i / 10
//count++
//}
//return uint64(count)
//}

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
