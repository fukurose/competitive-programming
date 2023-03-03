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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func main() {
	m := scanInt()
	n := scanInt()
	a := scanInt()

	prices := make([][]int, m)
	for i := range prices {
		prices[i] = make([]int, n)
		for j := range prices[i] {
			prices[i][j] = scanInt()
		}
	}

	minPrices := make([]int, n)
	dp := make([][]int, n)
	for i := range minPrices {
		dp[i] = make([]int, m)
		minPrices[i] = 9999999
	}

	for i := range dp[0] {
		dp[0][i] = prices[i][0]
	}

	for i := 0; i < n; i++ {
		for j := range dp[i] {
			price := prices[j][i]
			if i > 0 {
				dp[i][j] = min(dp[i-1][j]+price-a, price+minPrices[i-1])
			}
			minPrices[i] = min(minPrices[i], dp[i][j])
		}
	}

	ans := dp[n-1][0]

	for _, v := range dp[n-1] {
		ans = min(ans, v)
	}

	fmt.Println(ans)
}
