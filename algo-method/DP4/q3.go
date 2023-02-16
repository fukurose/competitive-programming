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
	m := scanInt()
	w := make([]int, n)
	for i := range w {
		w[i] = scanInt()
	}

	dp := make([][]bool, n)
	judge := make([][]int, n)
	for i := range dp {
		dp[i] = make([]bool, m+1)
		judge[i] = make([]int, m+1)
		for j := range dp[i] {
			judge[i][j] = 0b10

		}
	}

	dp[0][0] = true
	if w[0] <= m {
		dp[0][w[0]] = true
		judge[0][w[0]] = 0b01
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if !dp[i-1][j] {
				continue
			}

			// ボールを取らない場合
			if dp[i][j] {
				judge[i][j] = judge[i][j] | judge[i-1][j]
			} else {
				judge[i][j] = judge[i-1][j]
			}
			dp[i][j] = true

			// ボールを取る場合
			if j+w[i] <= m {
				prev := judge[i-1][j]
				now := prev ^ 0b11
				if prev == 0b11 {
					now = 0b11
				}
				if dp[i][j+w[i]] {
					judge[i][j+w[i]] = judge[i][j+w[i]] | now
				} else {
					judge[i][j+w[i]] = now
				}
				dp[i][j+w[i]] = true
			}
		}
	}

	if judge[n-1][m]&(1<<0) != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
