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
	m := scanInt()
	liked := make([]int, m+1)

	for i := 0; i < n; i++ {
		foods := scanInt()
		for j := 0; j < foods; j++ {
			liked[scanInt()]++
		}
	}

	count := 0
	for _, v := range liked {
		if v == n {
			count++
		}
	}

	fmt.Println(count)
}
