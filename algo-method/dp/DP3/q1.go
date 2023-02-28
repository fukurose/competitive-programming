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
		dp[i] = make([]bool, m)
	}

	arr := make([]int, n-1)

	for i := range arr {
		arr[i] = scanInt()
	}

	dp[0][0] = true

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			leftMove := arr[i-1]

			if dp[i-1][j] {
				dp[i][j] = true
			} else if j-leftMove >= 0 && dp[i-1][j-leftMove] {
				dp[i][j] = true
			}
		}
	}

	ans := 0
	for i := range dp[n-1] {
		if dp[n-1][i] {
			ans++
		}

	}
	fmt.Println(ans)
}
