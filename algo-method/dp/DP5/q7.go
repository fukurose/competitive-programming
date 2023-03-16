package main

import (
	"bufio"
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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func isOK(m int, p int, bit1 int, bit2 int) bool {
	for j := 0; j < m; j++ {
		if bit1 > 0&1<<j {
			continue
		}
		cnt_white := 0

	}

}

func main() {
	m := scanInt()
	n := scanInt()

	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
		for j := range arr[i] {
			arr[i][j] = scanInt()
		}
	}

	var inf = 1 << 30
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 1<<m)
		for j := range dp[i] {
			dp[i][j] = make([]int, 1<<m)
			for k := range dp[i] {
				dp[i][j][k] = inf
			}
		}
	}

	sumCol := make([][]int, n)
	for i := range sumCol {
		sumCol[i] = make([]int, 1<<m)
	}

	for i := 0; i < n; i++ {
		for p := 0; p < 1<<m; p++ {
			sum := 0
			for j := 0; j < m; j++ {
				if p > 0&1<<j {
					sum += arr[j][i]
				}
			}
			sumCol[i][p] = sum
		}
	}
	full := (1 << m) - 1
	dp[0][full][full] = 0
	for i := 0; i < n; i++ {
		for p := 0; p < 1<<m; p++ {
			for bit1 := 0; bit1 < 1<<m; bit1++ {
				for bit2 := 0; bit2 < 1<<m; bit2++ {
					if isOK(m, p, bit1, bit2) {
						dp[i+1][p][bit1] = min(dp[i+1][p][bit1], dp[i][bit1][bit2]+sumCol[i][p])
					}
				}
			}
		}
	}

}
