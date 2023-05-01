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

func main() {
	cube := scanString()

	zeroCount := 0
	oneCount := 0

	for _, r := range cube {
		if string(r) == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}
	deleteCount := min(zeroCount, oneCount)
	fmt.Println(deleteCount * 2)

}
