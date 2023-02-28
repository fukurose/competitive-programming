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

func chMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	n := scanInt()
	a := make([]int, n)
	b := make([]int, n)

	for i := range a {
		a[i] = scanInt()
	}

	for i := range b {
		b[i] = scanInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = a[0]
	dp[0][1] = b[0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + a[i]
		dp[i][1] = dp[i-1][1] + b[i]

		if dp[i][0] > dp[i][1] {
			dp[i][0] = chMin(dp[i][0], dp[i][1]+a[i])
		} else {
			dp[i][1] = chMin(dp[i][1], dp[i][0]+b[i])
		}
	}

	ans := chMin(dp[n-1][0], dp[n-1][1])
	fmt.Println(ans)
}
