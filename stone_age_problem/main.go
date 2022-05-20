//https://codeforces.com/problemset/problem/1679/B
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(in *bufio.Reader) [][]uint {
	var input [][]uint
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

func AAtoi(st string) []uint {
	var result []uint
	array := strings.Fields(st)
	for _, s := range array {
		i, _ := strconv.Atoi(s)
		result = append(result, uint(i))
	}
	return result
}

func sumOnce(array []uint) uint {
	result := uint(0)
	for _, v := range array {
		result += v
	}
	return result
}

func q1(array []uint, pos uint, replacement uint) []uint {
	array[pos] = replacement
	return array
}

func q2(array []uint, replacement uint) []uint {
	for i := range array {
		array[i] = replacement
	}
	return array
}

func contains(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sumPos(a []uint, positions []uint) uint {
	result := uint(0)
	for _, p := range positions {
		result += a[p]
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	input := getInput(in)
	n := input[0][0]
	//q := input[0][1]
	a := input[1]
	lq := ""
	i := uint(0)
	var positions []uint
	sum := sumOnce(a)
	for _, query := range input[2:] {
		if query[0] == 1 {
			pos := uint(query[1])
			rep := uint(query[2])
			//Actualizo el array y guardo donde actualicé
			q1(a, pos-1, rep)
			if !contains(positions, pos-1) {
				positions = append(positions, pos-1)
			}
			// Analizo casos previos
			if lq == "q2" {
				sum = sum - i + rep
			} else if lq == "q1" {
				sum = i*(uint(n)-uint(len(positions))) + sumPos(a, positions)
			} else {
				// Caso inicial
				sum = sumOnce(a)
			}
			lq = "q1"
		}
		if query[0] == 2 {
			rep := query[1]
			sum = n * rep
			i = rep
			lq = "q2"
			positions = []uint{}
		}
		fmt.Println(sum)
	}
}
