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
	m := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] += 1
	if arr[0] <= m {
		dp[0][arr[0]] += 1
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] < 1 {
				continue
			}

			dp[i][j] += dp[i-1][j]
			dp[i][j] %= 1000

			if j+arr[i] <= m {
				dp[i][j+arr[i]] += dp[i-1][j]
				dp[i][j+arr[i]] %= 1000
			}
		}
	}

	fmt.Println(dp[n-1][m])
}
