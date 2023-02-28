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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()
	dp := make([][3]int, n)
	price := make([][3]int, n)

	for i := range price {
		for j := 0; j < 3; j++ {
			price[i][j] = scanInt()
		}
	}

	dp[0] = price[0]

	for i := 1; i < n; i++ {
		dp[i][0] = price[i][0] + max(dp[i-1][1], dp[i-1][2])
		dp[i][1] = price[i][1] + max(dp[i-1][0], dp[i-1][2])
		dp[i][2] = price[i][2] + max(dp[i-1][0], dp[i-1][1])
	}

	ans := -1
	for _, v := range dp[n-1] {
		ans = max(ans, v)
	}

	fmt.Println(ans)
}
