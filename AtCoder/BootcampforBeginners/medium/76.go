package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	n := scanInt()
	s := scanString()

	left := make(map[string]int, 2)
	right := make(map[string]int, 2)

	arr := strings.Split(s, "")
	for _, c := range arr {
		right[c]++
	}

	ans := n

	for _, c := range arr {
		right[c]--
		ans = min(ans, left["W"]+right["E"])
		left[c]++
	}
	fmt.Println(ans)
}
