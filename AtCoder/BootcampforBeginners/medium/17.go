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
	t := scanInt()
	times := make([]int, n)

	for i := range times {
		times[i] = scanInt()
	}

	total := 0
	for i := 0; i < n-1; i++ {
		if times[i]+t > times[i+1] {
			total += times[i+1] - times[i]
		} else {
			total += t
		}
	}

	// 最後の人が押した分は t 秒出続ける
	total += t

	fmt.Println(total)
}
