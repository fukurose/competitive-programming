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
	points := make([]int, n)

	total := 0
	for i := range points {
		points[i] = scanInt()
		total += points[i]
	}

	sort.Ints(points)
	if total%10 == 0 {
		for _, point := range points {
			if point%10 != 0 {
				total -= point
				break
			}
		}
	}

	if total%10 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(total)
	}
}
