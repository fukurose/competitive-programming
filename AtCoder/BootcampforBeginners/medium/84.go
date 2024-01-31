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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	sort.Ints(arr)
	ans := 0
	for i := 0; i < n; i++ {
		a := arr[i]
		for j := i + 1; j < n; j++ {
			b := arr[j]
			border := sort.Search(len(arr), func(i int) bool { return arr[i] >= a+b }) - 1
			ans += max(border-j, 0)
		}
	}
	fmt.Println(ans)
}
