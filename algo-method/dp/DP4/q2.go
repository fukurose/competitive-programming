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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

var INF = 99999999999

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	a := scanInt()
	b := scanInt()

	p := make([]int, n)
	for i := range p {
		p[i] = scanInt()
	}

	q := make([]int, n)
	for i := range q {
		q[i] = scanInt()
	}

	r := make([]int, n)
	for i := range r {
		r[i] = scanInt()
	}

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, 3)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[0][0] = p[0]
	dp[0][1] = q[0]
	dp[0][2] = r[0]

	for i := 0; i < n-1; i++ {
		// Pで購入する場合
		dp[i+1][0] = min(dp[i+1][0], dp[i][0]+p[i+1]-a)
		// 3日連続購入
		if i+2 < n {
			dp[i+2][0] = min(dp[i+2][0], dp[i+1][0]+p[i+2]-b)
		}
		dp[i+1][0] = min(dp[i+1][0], dp[i][2]+p[i+1])
		dp[i+1][0] = min(dp[i+1][0], dp[i][1]+p[i+1])

		// Qで購入する場合
		dp[i+1][1] = min(dp[i+1][1], dp[i][1]+q[i+1]-a)
		// 3日連続購入
		if i+2 < n {
			dp[i+2][1] = min(dp[i+2][1], dp[i+1][1]+q[i+2]-b)
		}
		dp[i+1][1] = min(dp[i+1][1], dp[i][0]+q[i+1])
		dp[i+1][1] = min(dp[i+1][1], dp[i][2]+q[i+1])

		// Rで購入する場合
		dp[i+1][2] = min(dp[i+1][2], dp[i][2]+r[i+1]-a)
		// 3日連続購入
		if i+2 < n {
			dp[i+2][2] = min(dp[i+2][2], dp[i+1][2]+r[i+2]-b)
		}
		dp[i+1][2] = min(dp[i+1][2], dp[i][0]+r[i+1])
		dp[i+1][2] = min(dp[i+1][2], dp[i][1]+r[i+1])
	}

	ans := INF
	for i := range dp[n-1] {
		ans = min(ans, dp[n-1][i])
	}
	fmt.Println(ans)
}
