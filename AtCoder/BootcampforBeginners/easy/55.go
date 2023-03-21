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

func isInt(a int) bool {
	for i := 1; i <= a; i++ {
		if i*i == a {
			return true
		} else if i*i > a {
			break
		}
	}
	return false
}

func calD(a, b []int, l int) int {
	var d float64 = 0
	for i := 0; i < l; i++ {
		d += math.Pow(float64(a[i]-b[i]), 2)
	}
	return int(d)
}

func main() {
	n := scanInt()
	d := scanInt()

	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, d)
		for j := range arr[i] {
			arr[i][j] = scanInt()
		}
	}

	ans := 0
	for i := range arr {
		for j := i + 1; j < n; j++ {
			d := calD(arr[i], arr[j], d)
			if isInt(d) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
