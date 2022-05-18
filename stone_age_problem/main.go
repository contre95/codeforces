//https://codeforces.com/problemset/problem/1679/B
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func AAtoi(array []string) []int {
	var result []int
	for _, s := range array {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func q1(array []int, pos int, replacement int) []int {
	array[pos] = replacement
	return array
}

func q2(array []int, replacement int) []int {
	for i := range array {
		array[i] = replacement
	}
	return array
}

func main() {
	input := getInput()
	a := AAtoi(strings.Split(input[1], " "))
	for _, query := range input[2:] {
		aquery := strings.Split(query, " ")
		if aquery[0] == "1" {
			pos, _ := strconv.Atoi(string(aquery[1]))
			rep, _ := strconv.Atoi(string(aquery[2]))
			q1(a, pos-1, rep)
			fmt.Println(sum(a))
		}
		if aquery[0] == "2" {
			rep, _ := strconv.Atoi(string(aquery[1]))
			q2(a, rep)
			fmt.Println(sum(a))
		}
	}
}
