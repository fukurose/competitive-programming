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

func keta(k int) int {
	if k == 0 {
		return 0
	}
	ans := 1
	for i := 1; i < k; i++ {
		ans *= 10
	}
	return ans
}

func main() {
	a := scanInt()
	b := scanInt()
	x := scanInt()

	maxK := 10
	for {
		n := keta(maxK)
		if (a*n)+(b*maxK) <= x || maxK == 0 {
			break
		} else {
			maxK--
		}
	}

	if maxK == 10 || maxK == 0 {
		fmt.Println(keta(maxK))
	} else {
		r := x - (b * maxK)
		ans := r / a
		if ans >= keta(maxK+1) {
			ans = keta(maxK+1) - 1
		}
		fmt.Println(ans)
	}
}
