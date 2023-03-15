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
	k := scanInt()
	q := scanInt()

	challengers := make([]int, n+1)
	for i := 0; i < q; i++ {
		a := scanInt()
		challengers[a]++
	}

	for i := 1; i < n+1; i++ {
		if q-challengers[i] < k {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
