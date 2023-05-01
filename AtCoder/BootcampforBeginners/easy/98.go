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

func main() {
	arr := make([]int, 3)
	total := 1
	for i := range arr {
		arr[i] = scanInt()
		total *= arr[i]
	}

	ans := 0
	if total%2 != 0 {
		sort.Ints(arr)
		ans = arr[0] * arr[1]
	}
	fmt.Println(ans)
}
