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
	counter := make(map[int]int, n)

	for i := 0; i < n; i++ {
		counter[scanInt()]++
	}

	ans := 0

	for n, c := range counter {
		if n < c {
			ans += c - n
		} else if n > c {
			ans += c
		}
	}
	fmt.Println(ans)
}
