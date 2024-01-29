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
	m := scanInt()

	var ans int

	if n-m > 1 || m-n > 1 {
		ans = 0
	} else {
		ans = 1
		for i := 1; i <= n; i++ {
			ans *= i
			ans %= 1000000007
		}

		for i := 1; i <= m; i++ {
			ans *= i
			ans %= 1000000007
		}
	}

	if n == m {
		ans *= 2
		ans %= 1000000007
	}

	fmt.Println(ans)
}
