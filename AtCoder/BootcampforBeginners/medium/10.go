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
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	n := scanInt()
	counter := make(map[int]int, 100000)

	for i := 0; i < n; i++ {
		i := scanInt()
		counter[i-1]++
		counter[i]++
		counter[i+1]++

	}

	ans := -1
	for _, count := range counter {
		ans = max(ans, count)
	}

	fmt.Println(ans)
}
