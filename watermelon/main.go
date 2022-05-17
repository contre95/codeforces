package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	for scanner.Scan() {
		input := scanner.Text()
		n, _ = strconv.Atoi(input)
	}
	if n%2 == 0 && n > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
