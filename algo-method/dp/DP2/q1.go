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
	dp := make([][4]int, 4)

	for i := range dp[0] {
		dp[0][i] = scanInt()
	}

	for i := 1; i < 4; i++ {
		for j := range dp[i] {
			sum := dp[i-1][j]

			if j-1 >= 0 {
				sum += dp[i-1][j-1]
			}

			if j+1 < 4 {
				sum += dp[i-1][j+1]
			}
			dp[i][j] = sum
		}
	}

	fmt.Println(dp[3][3])
}
