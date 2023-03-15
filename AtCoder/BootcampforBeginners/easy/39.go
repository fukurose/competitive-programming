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
	max := make([]int, 2)
	for i := range arr {
		arr[i] = scanInt()
		max = append(max, arr[i])
		sort.Ints(max)
		max = max[1:]
	}

	for i := range arr {
		if arr[i] == max[1] {
			fmt.Println(max[0])
		} else {
			fmt.Println(max[1])
		}
	}
}
