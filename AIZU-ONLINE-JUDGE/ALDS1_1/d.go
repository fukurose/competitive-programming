package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	input := scanner.Text()
	i, _ := strconv.Atoi(input)
	return i
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var len int
	len = scanInt()

	max := -9999999999
	min := scanInt()
	for i := 1; i < len; i++ {
		num := scanInt()
		max = Max(max, num-min)
		min = Min(min, num)
	}

	fmt.Println(max)
}
