package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var INF = 99999

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func chMin(a, b int) int {
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
	m := scanInt()
	k := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0
	if arr[0] <= m {
		dp[0][arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] == INF {
				continue
			}

			dp[i][j] = chMin(dp[i][j], dp[i-1][j])

			if j+arr[i] <= m {
				dp[i][j+arr[i]] = chMin(dp[i][j+arr[i]], dp[i-1][j]+1)
			}
		}
	}

	if dp[n-1][m] <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
