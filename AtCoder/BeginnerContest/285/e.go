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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	days := scanInt()
	points := make([]int, days)
	pointSum := make([]int, days+1)

	for i := range points {
		points[i] = scanInt()
		pointSum[i+1] = pointSum[i] + points[i]
	}

	dp := make([]int, days) //　必ずひとつは休日なので、days の配列として、 dp[i] は i日目の最大値とする

	dp[0] = 0

	for d := 1; d < days; d++ {
		// d日間をどれも休みにしない場合
		point := pointSum[d/2] * 2
		if d%2 != 0 {
			point += points[(d / 2)]
		}
		dp[d] = point

		// 休みにする場合は、どの日を休みにするか。
		for h := 1; h <= d; h++ {
			// 休みの左側と右側を足した値
			dp[d] = max(dp[d], dp[h-1]+dp[d-h])
		}
	}

	fmt.Println(dp[days-1])
}
