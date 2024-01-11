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
	arr := make([]int, n)
	table := make(map[int]int)

	for i := range arr {
		arr[i] = scanInt()
		table[arr[i]]++
	}
	longSticks := make([]int, 3)

	for n, c := range table {
		if c >= 2 {
			longSticks[0] = n
			sort.Ints(longSticks)
		}
	}

	long1 := longSticks[2]
	long2 := longSticks[1]

	var ans int
	if table[long1] >= 4 {
		ans = long1 * long1
	} else {
		ans = long1 * long2
	}

	fmt.Println(ans)
}
