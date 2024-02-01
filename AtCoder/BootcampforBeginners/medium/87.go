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

	arrA := make([]int, n)
	arrB := make([]int, n)
	diff := make([]int, n)

	for i := range arrA {
		arrA[i] = scanInt()
	}

	for i := range arrB {
		arrB[i] = scanInt()
		diff[i] = arrA[i] - arrB[i]
	}

	sort.Ints(diff)

	lIndex := 0
	rIndex := n - 1

	for diff[lIndex] < 0 {
		for diff[lIndex] < 0 && lIndex < rIndex {
			if diff[rIndex]+diff[lIndex] > 0 {
				diff[rIndex] += diff[lIndex]
				diff[lIndex] = 0
			} else {
				diff[lIndex] += diff[rIndex]
				rIndex--
			}
		}
		lIndex++
	}

	sort.Ints(diff)
	if lIndex == 0 {
		fmt.Println(0)
	} else if diff[0] < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(lIndex + (n - rIndex))
	}

}
