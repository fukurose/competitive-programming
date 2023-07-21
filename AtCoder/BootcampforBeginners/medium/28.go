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
	sum := make([]float64, 1001)
	sum[0] = 0
	for i := 1; i <= 1000; i++ {
		sum[i] = float64(i) + sum[i-1]
	}

	n := scanInt()
	k := scanInt()
	p := make([]int, n)

	for i := range p {
		p[i] = scanInt()
	}

	sumExpectedValues := make([]float64, n)
	sumExpectedValues[0] = sum[p[0]] / float64(p[0])
	for i := 1; i < n; i++ {
		sumExpectedValues[i] = (sum[p[i]] / float64(p[i])) + sumExpectedValues[i-1]
	}

	var max float64 = sumExpectedValues[k-1]
	for i := k; i < n; i++ {
		v := sumExpectedValues[i] - sumExpectedValues[i-k]
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
}
