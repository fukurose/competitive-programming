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

func scanFloat() float64 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return float64(n)
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
	t := scanFloat()
	a := scanFloat()

	temperatures := make([]float64, n)
	for i := range temperatures {
		x := scanFloat()
		temperatures[i] = t - x*0.006
	}

	ans := 1
	min_diff := math.Abs(a - temperatures[0])

	for i := range temperatures {
		diff := math.Abs(a - temperatures[i])
		if min_diff > diff {
			ans = i + 1
			min_diff = diff
		}
	}

	fmt.Println(ans)
}
