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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func main() {
	n := scanInt()

	before := make([]int, n+1)
	after := make([]int, n+1)
	b := make([]int, n)

	for i := range before {
		before[i] = scanInt()
		after[i] = before[i]
	}

	for i := range b {
		b[i] = scanInt()
	}

	for i := range b {
		if after[i] >= b[i] {
			after[i] -= b[i]
		} else {
			after[i+1] -= (b[i] - after[i])
			after[i] = 0

			after[i+1] = max(after[i+1], 0)
		}
	}

	ans := 0
	for i := range before {
		ans += before[i] - after[i]
	}
	fmt.Println(ans)
}
