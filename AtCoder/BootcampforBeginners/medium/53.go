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

func calPrice(a, b int) int {
	p := a - b
	if p < 0 {
		return p * -1
	} else {
		return p
	}

}

func main() {
	n := scanInt()
	arr := make([]int, n+2)
	for i := 1; i < n+1; i++ {
		arr[i] = scanInt()
	}

	sum := make([]int, n+2)
	for i := 1; i < n+2; i++ {
		sum[i] = sum[i-1] + calPrice(arr[i-1], arr[i])
	}

	for i := 1; i < n+1; i++ {
		ans := sum[n+1] - sum[i+1]
		ans += calPrice(arr[i-1], arr[i+1]) + sum[i-1]
		fmt.Println(ans)
	}
}
