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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func counter(s string, target rune) int {
	start := 0
	count := -1
	for i, r := range s {
		if r == target {
			count = max(i-start, count)
			start = i + 1
		}
	}

	count = max(len(s)-start, count)

	return count
}

func main() {
	s := scanString()

	ans := 100
	// r is 'a'
	var r rune = 97
	for i := 0; i < 26; i++ {
		c := counter(s, r)
		ans = min(ans, c)
		r = r + 1
	}

	fmt.Println(ans)
}
