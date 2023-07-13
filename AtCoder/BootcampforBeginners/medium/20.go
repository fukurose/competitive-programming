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
	toggle := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		a := scanInt()
		toggle[a] = !toggle[a]
	}

	count := 0
	for _, v := range toggle {
		if v {
			count++
		}
	}

	fmt.Println(count)
}
