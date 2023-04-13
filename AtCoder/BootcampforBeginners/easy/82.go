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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	q := scanInt()
	h := scanInt()
	s := scanInt()
	d := scanInt()

	n := scanInt()

	miniPriceBy1L := 4 * q
	miniPriceBy1L = min(miniPriceBy1L, 2*h)
	miniPriceBy1L = min(miniPriceBy1L, s)

	miniPriceBy2L := min(miniPriceBy1L*2, d)

	ans := 0
	if n >= 2 {
		ans += (n / 2) * miniPriceBy2L
		n %= 2
	}
	if n > 0 {
		ans += miniPriceBy1L
	}
	fmt.Println(ans)

}
