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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func Atoi(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func main() {
	s := scanString()
	l := len(s)
	a := scanInt()

	dp := make([][][]int, l)
	for i := range dp {
		dp[i] = make([][]int, a)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	first := Atoi(string(s[0]))
	for i := 0; i < 10; i++ {
		if i < first {
			dp[0][i%a][1]++
		} else if i == first {
			dp[0][i%a][0]++
		}
	}

	for i := 1; i < l; i++ {
		n := Atoi(string(s[i]))
		for j := 0; j < a; j++ {
			for k := 0; k < 10; k++ {
				if k < n {
					dp[i][(j+k)%a][1] += (dp[i-1][j][0] + dp[i-1][j][1])
				} else if k == n {
					dp[i][(j+k)%a][0] += dp[i-1][j][0]
					dp[i][(j+k)%a][1] += dp[i-1][j][1]
				} else {
					dp[i][(j+k)%a][1] += dp[i-1][j][1]
				}
			}
		}
	}

	fmt.Println(dp[l-1][0][0] + dp[l-1][0][1] - 1)
}
