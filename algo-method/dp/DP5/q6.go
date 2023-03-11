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
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func init() {
	scanner.Split(bufio.ScanWords)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isOK(x, n, k int, sum []int) bool {
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if sum[i+1]-sum[j] >= x {
				if j != 0 && dp[j] == 0 {
					continue
				}
				dp[i+1] = max(dp[i+1], dp[j]+1)
			}
		}
	}

	return dp[n] >= k

}

func main() {
	n := scanInt()
	k := scanInt()
	sum := make([]int, n+1)
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
		sum[i+1] = sum[i] + a[i]
	}

	start := 0
	end := 99999999
	mid := (start + end) / 2

	for math.Abs(float64(start-end)) > 1 {
		if isOK(mid, n, k, sum) {
			start = mid
		} else {
			end = mid
		}
		mid = (start + end) / 2
	}
	fmt.Println(mid)
}
