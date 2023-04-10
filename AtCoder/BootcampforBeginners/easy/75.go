package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	n := scanInt()
	k := scanInt()

	h_arr := make([]int, n)
	for i := range h_arr {
		h_arr[i] = scanInt()
	}

	sort.Ints(h_arr)

	ans := 999999999
	for i := k - 1; i < n; i++ {
		ans = min(ans, h_arr[i]-h_arr[i-(k-1)])
	}

	fmt.Println(ans)
}
