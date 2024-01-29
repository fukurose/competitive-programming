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

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	ans := gcd(a, b)
	ans = a / ans
	return ans * b
}

func main() {
	n := scanInt()
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}
	ans := arr[0]

	for i := 1; i < n; i++ {
		ans = lcm(ans, arr[i])
	}

	fmt.Println(ans)
}
