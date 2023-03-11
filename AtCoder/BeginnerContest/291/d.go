package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var limit int = 998244353

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

func main() {
	n := scanInt()
	cards := make([][]int, n)
	dp := make([][]int, n)

	for i := range cards {
		// 表と裏
		cards[i] = make([]int, 2)
		dp[i] = make([]int, 2)

		for j := range cards[i] {
			cards[i][j] = scanInt()
		}
	}

	// 一枚では重複することはないのでカウントアップ
	dp[0][0]++
	dp[0][1]++

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			// ひとつ前の表と一致するか
			if cards[i][j] != cards[i-1][0] {
				dp[i][j] += dp[i-1][0]
			}

			// ひとつ前の裏と一致するか
			if cards[i][j] != cards[i-1][1] {
				dp[i][j] += dp[i-1][1]
			}

			dp[i][j] %= limit
		}
	}

	ans := (dp[n-1][0] + dp[n-1][1]) % limit
	fmt.Println(ans)
}
