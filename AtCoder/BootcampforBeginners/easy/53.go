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
	load_counter := make([]int, n+1)
	m := scanInt()

	for i := 0; i < m; i++ {
		load_counter[scanInt()]++
		load_counter[scanInt()]++
	}

	for i := 1; i <= n; i++ {
		fmt.Println(load_counter[i])
	}
}
