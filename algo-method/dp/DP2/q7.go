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

func scanText() string {
	scanner.Scan()
	return scanner.Text()
}

func min(a, b int) int {
	if a < b {
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
	dp := make([][]int, n)
	point := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
		point[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 99999999999
			point[i][j] = scanInt()
		}
	}

	dp[0][n-1] = point[0][n-1]

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if i > 0 {
				dp[i][j] = min(dp[i][j], point[i][j]+dp[i-1][j])
			}
			if j < n-1 {
				dp[i][j] = min(dp[i][j], point[i][j]+dp[i][j+1])
			}
		}
	}

	fmt.Println(dp[n-1][0])
}
