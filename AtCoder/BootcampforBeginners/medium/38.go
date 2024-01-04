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
	table := make(map[int]int, n)
	for i := 0; i < n; i++ {
		n := scanInt()
		table[n]++
	}

	m := scanInt()
	isOK := true
	for i := 0; i < m; i++ {
		n := scanInt()
		if table[n] <= 0 {
			isOK = false
			break
		}
		table[n]--
	}

	if isOK {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
