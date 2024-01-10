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
	arr := make([]int, n)
	ans := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	index := 0
	for i := n - 1; i >= 0; i -= 2 {
		ans[index] = arr[i]
		index++
	}

	for i := n % 2; i < n; i += 2 {
		ans[index] = arr[i]
		index++
	}

	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", ans[i])
	}

	fmt.Println(ans[n-1])
}
