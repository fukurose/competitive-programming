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
	table := make(map[string]int, n)
	max_occurrence := 0
	for i := 0; i < n; i++ {
		s := scanString()
		table[s]++
		max_occurrence = max(max_occurrence, table[s])
	}

	var ans []string

	for word, count := range table {
		if count == max_occurrence {
			ans = append(ans, word)
		}
	}

	sort.Strings(ans)
	for _, word := range ans {
		fmt.Println(word)
	}

}
