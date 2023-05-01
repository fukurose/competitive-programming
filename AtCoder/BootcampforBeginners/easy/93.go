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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	size := scanInt()
	arr := make([]int, size)

	leftSum := 0
	rightSum := 0
	for i := range arr {
		arr[i] = scanInt()
		rightSum += arr[i]
	}

	d := rightSum
	for _, v := range arr {
		leftSum += v
		rightSum -= v

		d = min(d, diff(leftSum, rightSum))
	}

	fmt.Println(d)
}
