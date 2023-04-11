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
	n := scanInt()
	points := make([]int, n*3)

	for i := range points {
		points[i] = scanInt()
	}

	sort.Ints(points)

	ans := 0
	for i := 1; i <= n; i++ {
		index := 3*n - (i * 2) // 2番目の最大値のため、2つ飛ばし
		ans += points[index]
	}

	fmt.Println(ans)
}
