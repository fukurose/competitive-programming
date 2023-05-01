package main

import (
	"bufio"
	"fmt"
	"math"
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

	for i := range arr {
		arr[i] = scanInt()
	}

	all := int(math.Pow(3, float64(n)))
	odd := 1
	for _, v := range arr {
		if v%2 == 0 {
			odd *= 2
		} else {
			odd *= 1
		}
	}

	fmt.Println(all - odd)
}
