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

func main() {
	n := scanInt()
	total := scanInt()

	ans := []int{-1, -1, -1}
	for x := 0; x <= n; x++ {
		for y := 0; y <= (n - x); y++ {
			r := total - (10000 * x) - (5000 * y)
			if r == 1000*(n-x-y) {
				ans[0] = x
				ans[1] = y
				ans[2] = n - x - y
				break
			}
		}
	}
	fmt.Println(ans[0], ans[1], ans[2])
}
