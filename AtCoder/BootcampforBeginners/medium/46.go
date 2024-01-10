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
	s := scanString()
	table := make(map[string]int, 3)

	for _, r := range s {
		table[string(r)]++
	}

	arr := []int{table["a"], table["b"], table["c"]}

	ans := "NO"
	if len(s) == 1 {
		ans = "YES"
	} else if len(s) == 2 {
		if table["a"] == 1 || table["b"] == 1 || table["c"] == 1 {
			ans = "YES"
		}
	} else {
		sort.Ints(arr)
		if arr[2]-arr[0] <= 1 {
			ans = "YES"
		}
	}

	fmt.Println(ans)
}
