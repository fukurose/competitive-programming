package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	dp := make([][]int, n)
	masu := make([][]string, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range masu {
		masu[i] = strings.Split(scanText(), "")
	}

	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && masu[i-1][j] == "." {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 && masu[i][j-1] == "." {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	fmt.Println(dp[n-1][n-1])
}
