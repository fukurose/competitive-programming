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

func diff(a, b int) int {
	ans := a - b
	if ans < 0 {
		ans *= -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	n := scanInt()
	sumArr := make([]int, n)

	sumArr[0] = scanInt()

	for i := 1; i < n; i++ {
		sumArr[i] = scanInt() + sumArr[i-1]
	}

	minD := diff(sumArr[0], sumArr[n-1]-sumArr[0])

	for i := 1; i < n-1; i++ {
		d := diff(sumArr[i], sumArr[n-1]-sumArr[i])
		minD = min(minD, d)
	}

	fmt.Println(minD)
}
