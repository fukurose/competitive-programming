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
	total_time := 0
	times := make([]int, n)
	for i := range times {
		times[i] = scanInt()
		total_time += times[i]
	}

	m := scanInt()
	ans := make([]int, m)

	for i := range ans {
		index := scanInt() - 1
		time := scanInt()

		diff := times[index] - time

		ans[i] = total_time - diff

	}

	for _, time := range ans {
		fmt.Println(time)
	}

}
