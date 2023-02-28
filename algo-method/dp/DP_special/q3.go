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

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	if arr[0] <= m {
		dp[0][arr[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if !dp[i-1][j] {
				continue
			}
			dp[i][j] = dp[i-1][j]

			if j+arr[i] <= m {
				dp[i][j+arr[i]] = true
			}
		}
	}

	if dp[n-1][m] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
