package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	size := scanInt()
	arr := make([]int, size)
	sorted := make([]int, size)
	for i := range arr {
		arr[i] = scanInt()
		sorted[i] = arr[i]
	}

	sort.Ints(sorted)

	center := sorted[size/2-1]

	for _, n := range arr {
		if n <= center {
			fmt.Println(sorted[size/2])
		} else {
			fmt.Println(center)
		}
	}
}
