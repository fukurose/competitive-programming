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
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func init() {
	scanner.Split(bufio.ScanWords)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()
	k := scanInt()

	dp := make([][]int, n)
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -999999999
		}
	}

	dp[0][0] = arr[0]
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		// arr[i]を消さない場合というのは、前回でも削除していない、もしくは削除し終わっている場合
		dp[i][0] = max(dp[i-1][0], dp[i-1][k]) + arr[i]

		for j := 1; j <= k; j++ {
			// arr[i]を消す場合は連続なので、前回でk-1消している場合
			dp[i][j] = dp[i-1][j-1]

			// 繰り返し削除可能のため、前回で k削除している分も考慮する
			if j == k {
				dp[i][j] = max(dp[i][j], dp[i-1][k])
			}
		}
	}

	ans := max(dp[n-1][0], dp[n-1][k])
	fmt.Println(ans)

}
