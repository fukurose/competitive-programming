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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func chMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	n := scanInt()
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	min := arr[0]

	count := 0
	for i := range arr {
		if min >= arr[i] {
			count++
		}

		min = chMin(min, arr[i])
	}

	fmt.Println(count)
}
