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
	arrA := make([]int, n)
	arrB := make([]int, n)

	for i := range arrA {
		arrA[i] = scanInt()
		arrB[i] = scanInt()
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		count := arrA[i] + ans
		if count%arrB[i] != 0 {
			diff := arrB[i] - count%arrB[i]
			ans += diff
		}
	}
	fmt.Println(ans)
}
