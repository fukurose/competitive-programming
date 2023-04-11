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
	heights := make([]int, n)

	for i := range heights {
		heights[i] = scanInt()
	}

	ans := "Yes"

	for i := 1; i < n; i++ {
		if heights[i-1] < heights[i] {
			heights[i]--
		}

		if heights[i-1] > heights[i] {
			ans = "No"
			break
		}
	}

	fmt.Println(ans)
}
