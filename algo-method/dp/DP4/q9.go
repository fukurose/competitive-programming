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

func main() {
	n := scanString()
	a := scanInt()

	size := len(n)
	nums := make([]int, size+1)
	for i, v := range n {
		num, _ := strconv.Atoi(string(v))
		nums[size-i] = num
	}

	dp := make([][]int, size+1)
	for i := range dp {
		dp[i] = make([]int, a)
	}

	for i := 1; i <= size; i++ {
		for j, v := range dp[i-1] {
			dp[i][j] = v
		}

		limit := 9
		if i == size {
			limit = nums[size]
		}

		for j := 1; j <= limit; j++ {
			dp[i][j%a]++
		}
	}
	fmt.Println(dp)
}
