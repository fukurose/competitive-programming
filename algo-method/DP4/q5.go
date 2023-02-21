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

var left int = 0
var right int = 1
var both int = 2

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	n := scanInt()

	points := make([][]int, n)
	dp := make([][]int, n)
	for i := range points {
		points[i] = make([]int, 2)
		dp[i] = make([]int, 3)
	}

	for i := range points {
		points[i][1] = scanInt()
	}

	for i := range points {
		points[i][0] = scanInt()
	}

	dp[0][left] = points[0][left]
	dp[0][right] = points[0][right]
	dp[0][both] = points[0][left] + points[0][right]

	for i := 1; i < n; i++ {
		dp[i][left] = min(dp[i-1][right], dp[i-1][both]) + points[i][left]
		dp[i][right] = min(dp[i-1][left], dp[i-1][both]) + points[i][right]
		dp[i][both] = min(min(dp[i-1][left], dp[i-1][right]), dp[i-1][both]) + points[i][right] + points[i][left]
	}

	fmt.Println(min(min(dp[n-1][left], dp[n-1][right]), dp[n-1][both]))

}
