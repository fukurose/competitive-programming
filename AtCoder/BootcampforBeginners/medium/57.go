package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	counter := make(map[string]int, n)

	for i := 0; i < n; i++ {
		s := scanString()
		arr := strings.Split(s, "")
		sort.Slice(arr, func(j, k int) bool {
			return arr[j] < arr[k]
		})
		sorted := strings.Join(arr, "")
		counter[sorted]++
	}

	ans := 0
	for _, count := range counter {
		ans += count * (count - 1) / 2
	}
	fmt.Println(ans)
}
