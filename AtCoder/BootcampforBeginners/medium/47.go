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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	p := scanInt()

	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 最初を選ばなかった場合
	dp[0][0]++
	// 最初を選んだ場合
	dp[0][arr[0]%2]++
	for i := 1; i < n; i++ {
		// 選ばなかった場合
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1]

		// 選んだ場合
		if arr[i]%2 == 0 {
			dp[i][0] += dp[i-1][0]
			dp[i][1] += dp[i-1][1]
		} else {
			dp[i][0] += dp[i-1][1]
			dp[i][1] += dp[i-1][0]
		}
	}

	fmt.Println(dp[n-1][p])
}
