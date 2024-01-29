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

const eps = 1e-14

func main() {
	a := scanInt()
	b := scanInt()
	c := scanInt()

	d := c - a - b
	if d > 0 && 4*a*b < d*d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
