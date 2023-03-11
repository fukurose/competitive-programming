package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var mod int = 1000000007

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
	k := scanInt()

	s := scanString()

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	dp[0][0] = 1
	for i, v := range s {
		num, _ := strconv.Atoi(string(v))

		for j := 0; j < k; j++ {
			// num を使わない場合
			dp[i+1][j] += dp[i][j]
			dp[i+1][j] %= mod

			// num を使う場合
			new_num := (j * 10) + num
			dp[i+1][new_num%k] += dp[i][j]
			dp[i+1][new_num%k] %= mod
		}

	}

	fmt.Println(dp[n][0] - 1)
}
