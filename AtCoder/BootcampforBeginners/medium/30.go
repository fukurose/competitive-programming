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

	ans := 0
	if n == 1 && m == 1 {
		ans = 1
	} else if n == 1 {
		ans = m - 2
	} else if m == 1 {
		ans = n - 2
	} else {
		// 端以外は全て裏
		ans = (n - 2) * (m - 2)
	}

	fmt.Println(ans)
}
