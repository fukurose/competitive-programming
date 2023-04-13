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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
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

	ans := n - 1
	for x := 2; x <= int(math.Sqrt(float64(n))); x++ {
		if n%x == 0 {
			y := n / x
			ans = min(ans, (x-1)+(y-1))
		}
	}
	fmt.Println(ans)
}
