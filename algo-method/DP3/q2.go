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

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	w := make([]int, n)

	for i := range w {
		w[i] = scanInt()
	}

	dp[0][0] = true // 取らない場合
	if w[0] <= m {
		dp[0][w[0]] = true // 取る場合
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			//取らない場合
			dp[i][j] = dp[i-1][j]

			if !dp[i][j] && j-w[i] >= 0 {
				//取る場合
				dp[i][j] = dp[i-1][j-w[i]]
			}
		}
	}

	if dp[n-1][m] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
