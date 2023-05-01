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
	arr := []int{5, 4, 3, 2, 1}

	x := scanInt()
	r := x % 100

	count := 0
	for _, v := range arr {
		if r >= v {
			q := r / v
			r -= q * v
			count += q
		}
	}

	if count*100 <= x {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
