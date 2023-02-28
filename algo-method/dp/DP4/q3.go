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

var even int = 0
var odd int = 1

func main() {
	n := scanInt()
	m := scanInt()
	w := make([]int, n)
	for i := range w {
		w[i] = scanInt()
	}

	dp := make([][][2]bool, n)
	for i := range dp {
		dp[i] = make([][2]bool, m+1)
		for j := range dp[i] {
			dp[i][j] = [2]bool{}
		}
	}

	// ボールを取らない場合
	dp[0][0][even] = true

	// ボールを取る場合
	if w[0] <= m {
		dp[0][w[0]][odd] = true
	}

	for i := 0; i < n-1; i++ {
		for j := range dp[i] {
			for k := range dp[i][j] {
				if !dp[i][j][k] {
					continue
				}

				// ボールを取らない場合
				dp[i+1][j][k] = true

				// ボールを取る場合
				if j+w[i+1] <= m {
					dp[i+1][j+w[i+1]][(k+1)%2] = true
				}
			}
		}
	}

	if dp[n-1][m][odd] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
