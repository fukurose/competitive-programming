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

func max(a, b int) int {
	if a > b {
		return a
	} else {

		return b
	}
}

func main() {
	n := scanInt()
	x := scanInt()
	dp := make([]int, n+1)
	sum := make([]int, n+1)
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
		sum[i+1] = sum[i] + a[i]
	}

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

	fmt.Println(dp[n])
}
