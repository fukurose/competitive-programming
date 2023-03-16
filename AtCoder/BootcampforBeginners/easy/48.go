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
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	count := 0
	ans := 0
	for i := 1; i < n; i++ {
		if arr[i] <= arr[i-1] {
			count++
			ans = max(ans, count)
		} else {
			count = 0
		}
	}

	fmt.Println(ans)
}
