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
	s := scanInt()
	c := scanInt()

	scc := 0

	if c >= s*2 {
		scc += s
		c -= s * 2
		s = 0
		if c >= 4 {
			scc += c / 4
			c = 0
		}
	} else {
		scc += c / 2
		s -= c / 2
		c = 0
	}

	fmt.Println(scc)
}
