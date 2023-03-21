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
	b := make([]int, n-1)
	a := make([]int, n)

	for i := range b {
		b[i] = scanInt()
		a[i] = b[i]
	}
	a[n-1] = b[n-2]

	for i := range b {
		if a[i] > b[i] {
			a[i] = b[i]
		}

		if a[i+1] > b[i] {
			a[i+1] = b[i]
		}
	}

	ans := 0
	for _, v := range a {
		ans += v
	}
	fmt.Println(ans)
}
