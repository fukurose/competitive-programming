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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sW := 0
	sH := 0
	eW := scanInt()
	eH := scanInt()
	n := scanInt()

	for i := 0; i < n; i++ {
		x := scanInt()
		y := scanInt()
		f := scanInt()

		switch f {
		case 1:
			sW = max(sW, x)
		case 2:
			eW = min(eW, x)
		case 3:
			sH = max(sH, y)
		case 4:
			eH = min(eH, y)
		}
	}

	w := max(eW-sW, 0)
	h := max(eH-sH, 0)
	fmt.Println(w * h)
}
