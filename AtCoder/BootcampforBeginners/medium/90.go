package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	q := scanInt()
	s := scanString()

	arr := strings.Split(s, "")
	counter := make([]int, n)

	sum := 0
	for i := 1; i < n; i++ {
		if arr[i-1] == "A" && arr[i] == "C" {
			sum++
		}
		counter[i] = sum
	}

	ans := make([]int, q)
	for i := range ans {
		l := scanInt()
		r := scanInt()
		ans[i] = counter[r-1] - counter[l-1]

	}
	for _, x := range ans {
		fmt.Println(x)

	}
}
