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

func main() {
	n := scanInt()
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[1][2] = 1
	dp[2][1] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] += dp[i-1][j]
			dp[i][j] += dp[i][j-1]
		}
	}

	fmt.Println(dp[n][n])
}
