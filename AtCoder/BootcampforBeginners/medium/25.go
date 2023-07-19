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
	k := scanInt()

	arr := make([]int, n)
	m := make(map[int]int, n)

	for i := range arr {
		arr[i] = scanInt()
		m[arr[i]]++
	}

	values := []int{}
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)

	needChangeCount := len(values) - k

	ans := 0

	for i := 0; i < needChangeCount; i++ {
		ans += values[i]
	}

	fmt.Println(ans)

}
