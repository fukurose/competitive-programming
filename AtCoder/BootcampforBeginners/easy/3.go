package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
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

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt() // n は不要
	a_limit := scanInt()
	b_limit := scanInt()
	limit := a_limit + b_limit

	s := scanString()

	results := make([]bool, n)

	a_passes := 0
	b_passes := 0
	for i, r := range s {
		passes := a_passes + b_passes
		if passes >= limit {
			break
		}
		switch string(r) {
		case "a":
			a_passes++
			results[i] = true
		case "b":
			if b_passes < b_limit {
				b_passes++
				results[i] = true
			}
		}
	}

	for _, result := range results {
		if result {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
