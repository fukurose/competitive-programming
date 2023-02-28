package main

import (
	"bufio"
	"fmt"
	"math"
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

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func main() {
	n := scanInt()
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)

	for i := range a {
		a[i] = scanInt()
	}

	for i := range b {
		b[i] = scanInt()
	}

	for i := range c {
		c[i] = scanInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 1; i < n; i++ {
		// a[i] までの最小値
		cost_from_a := abs(a[i-1]-a[i]) + dp[i-1][0]
		cost_from_b := abs(b[i-1]-a[i]) + dp[i-1][1]
		cost_from_c := abs(c[i-1]-a[i]) + dp[i-1][2]

		dp[i][0] = chMin(cost_from_a, chMin(cost_from_b, cost_from_c))

		// b[i] までの最小値
		cost_from_a = abs(a[i-1]-b[i]) + dp[i-1][0]
		cost_from_b = abs(b[i-1]-b[i]) + dp[i-1][1]
		cost_from_c = abs(c[i-1]-b[i]) + dp[i-1][2]

		dp[i][1] = chMin(cost_from_a, chMin(cost_from_b, cost_from_c))

		// c[i] までの最小値
		cost_from_a = abs(a[i-1]-c[i]) + dp[i-1][0]
		cost_from_b = abs(b[i-1]-c[i]) + dp[i-1][1]
		cost_from_c = abs(c[i-1]-c[i]) + dp[i-1][2]

		dp[i][2] = chMin(cost_from_a, chMin(cost_from_b, cost_from_c))
	}

	ans := dp[n-1][0]
	for i := range dp[n-1] {
		ans = chMin(ans, dp[n-1][i])
	}

	fmt.Println(ans)
}
