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
	a := scanInt()
	b := scanInt()

	ans := 0

	if n >= 2 {
		if b >= a {
			min := a + b + (n-2)*a
			max := a + b + (n-2)*b
			ans = max - min + 1
		}
	} else if n == 1 {
		if a == b {
			ans = 1
		}
	}

	fmt.Println(ans)
}
