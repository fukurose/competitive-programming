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

func f(m int, arr []int) int {

	ans := 0
	for i := range arr {
		ans += m % arr[i]
	}
	return ans
}

func main() {
	n := scanInt()
	arr := make([]int, n)

	m := 1
	for i := range arr {
		arr[i] = scanInt()
		m += arr[i]
	}

	fmt.Println(f(m, arr))
}
