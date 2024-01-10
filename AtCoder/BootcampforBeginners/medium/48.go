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

func min(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	w := scanInt()
	h := scanInt()
	x := scanInt()
	y := scanInt()

	other := 0
	area := float64(w*h) / 2
	// 与えられた点が対角線の交点の場合
	if x*2 == w && y*2 == h {
		other = 1
	}

	fmt.Printf("%f %d\n", area, other)
}
