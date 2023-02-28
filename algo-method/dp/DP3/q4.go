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
		dp[i] = make([]int, m+1)
	}

	w := make([]int, n)
	v := make([]int, n)

	for i := range w {
		w[i] = scanInt()
	}

	for i := range v {
		v[i] = scanInt()
	}

	if w[0] < m {
		dp[0][w[0]] = v[0]
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			dp[i][j] = dp[i-1][j]

			if j-w[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}

	ans := 0
	for i := range dp[n-1] {
		ans = max(ans, dp[n-1][i])

	}
	fmt.Println(ans)
}
