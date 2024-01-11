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
	arr := make([][]int, n)

	for i := range arr {
		arr[i] = make([]int, 2)
		arr[i][0] = scanInt()
		arr[i][1] = scanInt()
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] <= arr[j][1]
	})

	time := 0
	enable := true

	for i := range arr {
		time += arr[i][0]

		if arr[i][1] < time {
			enable = false
			break
		}
	}

	if enable {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
