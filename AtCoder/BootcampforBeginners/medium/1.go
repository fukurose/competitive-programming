package main

import (
	"bufio"
	"fmt"
	"os"
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
	done := make([]bool, n)

	for i := range arr {
		arr[i] = scanInt() - 1
	}

	count := 1
	index := arr[0]
	done[0] = true

	for index != 1 {
		if done[index] {
			count = -1
			break
		}
		done[index] = true

		index = arr[index]
		count++
	}

	fmt.Println(count)
}
