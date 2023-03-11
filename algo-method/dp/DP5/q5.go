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
	n := scanInt()
	x := scanInt()

	arr := make([]int, n)
	sum := 0
	for i := range arr {
		arr[i] = scanInt()
		sum += arr[i]
	}

	dp := make([]int, n)

	// 少なくとも1つは保証されている
	ans := 1
	for {
		index := 0
		max_right := 0

		for i := 1; i < len(arr); i++ {
			left := dp[i-1] + arr[i-1]
			right := sum - left
			dp[i] = left

			if left >= x && right >= x {
				if max_right < right {
					index = i
					max_right = right
				}
			}
		}

		if index < 1 {
			break
		}
		arr = arr[index:]
		ans++
		sum -= dp[index]
		dp = make([]int, n)
	}

	fmt.Println(ans)
}
