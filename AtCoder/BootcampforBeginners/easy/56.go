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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func main() {
	a := scanInt()
	b := scanInt()
	k := scanInt()

	a_from := a
	a_to := min(a+(k-1), b)

	b_from := max(b-(k-1), a)
	b_to := b

	if a_to >= b_from {
		a_to = b_from - 1
	}

	for a := a_from; a <= a_to; a++ {
		fmt.Println(a)
	}

	for b := b_from; b <= b_to; b++ {
		fmt.Println(b)
	}
}
