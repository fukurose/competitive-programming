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
	n := scanInt()

	dp := make([][]bool, n)

	w := make([]int, n)

	w_sum := 0
	for i := range w {
		w[i] = scanInt()
		w_sum += w[i]
	}

	for i := range dp {
		dp[i] = make([]bool, w_sum+1)
	}

	dp[0][w[0]] = true

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] {
				if j+w[i] <= w_sum {
					dp[i][j+w[i]] = true
				}

				if w[i]-j >= 0 {
					dp[i][w[i]-j] = true
				}

				if j-w[i] >= 0 {
					dp[i][j-w[i]] = true
				}
			}
		}
	}

	for i := range dp[n-1] {
		if dp[n-1][i] {
			fmt.Println(i)
			break
		}
	}
}
