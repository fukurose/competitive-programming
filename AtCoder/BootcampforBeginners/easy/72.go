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

func isPowOf(x, y int) bool {
	for i := 2; ; i++ {
		p := int(math.Pow(float64(x), float64(i)))
		if p == y {
			return true
		} else if p > y {
			return false
		}
	}
}

func isPow(x int) bool {
	if x == 1 {
		return true
	}

	for i := 2; i < x; i++ {
		if isPowOf(i, x) {
			return true
		}
	}

	return false
}

func main() {
	x := scanInt()
	for {
		if isPow(x) {
			fmt.Println(x)
			break
		}
		x--
	}
}
