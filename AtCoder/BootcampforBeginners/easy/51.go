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

func sum(n int) int {
	s := strconv.Itoa(n)

	sum := 0
	for _, r := range s {
		num, _ := strconv.Atoi(string(r))
		sum += num
	}
	return sum
}

func main() {
	n := scanInt()
	a := scanInt()
	b := scanInt()

	ans := 0
	for i := 1; i <= n; i++ {
		num := sum(i)
		if a <= num && b >= num {
			ans += i

		}
	}

	fmt.Println(ans)
}
