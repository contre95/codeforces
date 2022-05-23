//https://codeforces.com/problemset/problem/1679/B
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(in *bufio.Reader) [][]uint64 {
	var input [][]uint64
	reader := bufio.NewReader(os.Stdin)
	line0, _ := reader.ReadString('\n')

	line0 = strings.ReplaceAll(line0, "\n", "")
	input = append(input, AAtoi(line0))
	lines_count, _ := strconv.Atoi(strings.Fields(line0)[1])
	for i := 0; i < lines_count+1; i++ {
		linen, _ := reader.ReadString('\n')
		linen = strings.ReplaceAll(linen, "\n", "")
		input = append(input, AAtoi(linen))
	}
	return input
}

func AAtoi(st string) []uint64 {
	var result []uint64
	array := strings.Fields(st)
	for _, s := range array {
		i, _ := strconv.Atoi(s)
		result = append(result, uint64(i))
	}
	return result
}

func sumOnce(array []uint64) uint64 {
	result := uint64(0)
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	input := getInput(in)
	n := input[0][0]
	a := input[1]
	i := uint64(0)
	positions := map[uint64]uint64{}
	sum := sumOnce(a)
	for _, query := range input[2:] {
		if query[0] == 1 {
			pos := uint64(query[1])
			rep := uint64(query[2])
			if _, ok := positions[pos-1]; ok {
				sum = sum + (rep - positions[pos-1])
			} else if i != 0 {
				sum = sum + rep - i
			} else {
				sum = sum + rep - a[pos-1]
			}
			positions[pos-1] = rep
		}
		if query[0] == 2 {
			rep := query[1]
			i = rep
			positions = map[uint64]uint64{}
			sum = n * i
		}
		fmt.Println(sum)
	}
}
