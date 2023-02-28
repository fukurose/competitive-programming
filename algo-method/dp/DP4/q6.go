package main

import (
	"bufio"
	"fmt"
	"math"
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

func pow2(x int) int {
	return int(math.Pow(2, float64(x)))
}

func isWhiteContinuous(x, size int) bool {
	for i := 0; i < size-1; i++ {
		if ((x>>i)&1 == 0) && ((x>>(i+1))&1 == 0) {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	m := scanInt()
	n := scanInt()

	points := make([][]int, m)
	dp := make([][]int, n)
	for i := range points {
		points[i] = make([]int, n)
		for j := range points[i] {
			points[i][j] = scanInt()
		}
	}

	for i := range dp {
		dp[i] = make([]int, pow2(m))
		for j := range dp[i] {
			dp[i][j] = 99999
		}
	}

	// 1列目を計算
	for b := 0; b < pow2(m); b++ {
		if isWhiteContinuous(b, m) {
			continue
		}

		point := 0
		for i := 0; i < m; i++ {
			if (b>>i)&1 == 1 {
				point += points[i][0]
			}
		}
		dp[0][b] = point
	}

	for i := 1; i < n; i++ {
		// 全パターンをbit検索する
		for b := 0; b < pow2(m); b++ {
			// 白が連続するパターンは除外
			if isWhiteContinuous(b, m) {
				continue
			}
			point := 0
			for k := 0; k < m; k++ {
				if (b>>k)&1 == 1 {
					point += points[k][i]
				}
			}

			// 1個前のパターンの組み合わせで白が隣り合わないものの最小値を設定
			for pre_b, pre_p := range dp[i-1] {
				if (b | pre_b) != pow2(m)-1 {
					continue
				}
				dp[i][b] = min(dp[i][b], pre_p+point)
			}
		}
	}

	ans := dp[n-1][0]
	for _, p := range dp[n-1] {
		ans = min(ans, p)
	}
	fmt.Println(ans)
}
