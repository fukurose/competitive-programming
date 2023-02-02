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
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = scanInt()
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := dp[i-1][j]

			if j-1 >= 0 {
				sum += dp[i-1][j-1]
			}

			if j+1 < n {
				sum += dp[i-1][j+1]
			}
			dp[i][j] = sum % 100
		}
	}

	fmt.Println(dp[n-1][n-1])
}
