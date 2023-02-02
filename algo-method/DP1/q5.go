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

func init() {
	scanner.Split(bufio.ScanWords)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()
	m := scanInt()

	arr := make([]int, n)
	ans := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + arr[i]
		loopLimit := min(i, m)
		for j := 2; j <= loopLimit; j++ {
			cost := ans[i-j] + j*arr[i]
			ans[i] = min(ans[i], cost)
		}
	}

	fmt.Println(ans[n-1])
}
