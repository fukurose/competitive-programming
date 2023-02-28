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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
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
		for j := range dp[i] {
			dp[i][j] = 9999
		}
	}

	w := make([]int, n)

	for i := range w {
		w[i] = scanInt()
	}

	dp[0][0] = 0
	if w[0] <= m {
		dp[0][w[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {

			dp[i][j] = dp[i-1][j]

			if j-w[i] >= 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-w[i]]+1)
			}
		}
	}

	ans := dp[n-1][m]
	if ans == 9999 {
		ans = -1
	}

	fmt.Println(ans)
}
