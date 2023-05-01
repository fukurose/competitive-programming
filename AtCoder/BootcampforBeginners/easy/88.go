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
	blueMap := make(map[string]int, n)

	for i := 0; i < n; i++ {
		blueMap[scanString()]++
	}

	m := scanInt()
	redMap := make(map[string]int, m)

	for i := 0; i < m; i++ {
		redMap[scanString()]++
	}

	ans := 0
	for k, v := range blueMap {
		ans = max(ans, v-redMap[k])
	}

	fmt.Println(ans)
}
