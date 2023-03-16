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
	s := scanString()

	oddZero := 0
	evenZero := 0
	for i, r := range s {
		if i%2 == 0 {
			if string(r) == "0" {
				oddZero++
			} else {
				evenZero++
			}
		} else {
			if string(r) == "1" {
				oddZero++
			} else {
				evenZero++
			}
		}
	}

	fmt.Println(min(oddZero, evenZero))
}
