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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	m := scanInt()

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	a := make([]int, n)
	b := make([]int, n)

	dp[0][0] = 0

	for i := 1; i < n; i++ {
		a[i] = scanInt()
	}

	for i := 1; i < n; i++ {
		b[i] = scanInt()
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] >= 0 {
				dp[i][j] = dp[i-1][j]
			}

			if j-a[i] >= 0 && dp[i-1][j-a[i]] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-a[i]]+b[i])
			}
		}
	}

	fmt.Println(dp[n-1][m-1])
}
