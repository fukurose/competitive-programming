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

func cal(a, b int) int {
	ans := a - b
	return ans * ans
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
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
	start := 101
	end := -101
	for i := range arr {
		arr[i] = scanInt()
		start = min(start, arr[i])
		end = max(end, arr[i])
	}

	ans := 1000000000
	for i := start; i <= end; i++ {
		cost := 0
		for _, n := range arr {
			cost += cal(i, n)
		}
		ans = min(ans, cost)
	}
	fmt.Println(ans)
}
