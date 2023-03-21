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
	x1 := scanInt()
	y1 := scanInt()
	x2 := scanInt()
	y2 := scanInt()

	diff_x := x2 - x1
	diff_y := y2 - y1

	x3 := x2 - diff_y
	y3 := y2 + diff_x

	x4 := x3 - diff_x
	y4 := y3 - diff_y

	fmt.Printf("%d %d %d %d \n", x3, y3, x4, y4)
}
