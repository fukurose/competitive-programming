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
	a := scanInt()
	b := scanInt()

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, a)
	}

	x := make([]int, n)

	for i := range x {
		x[i] = scanInt()
	}

	dp[0][0] = true
	dp[0][x[0]%a] = true

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
				dp[i][(j+x[i])%a] = true
			}
		}
	}

	if dp[n-1][b] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
