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
	s := scanString()
	max := len(s) - 1

	ans := 0
	for i, r := range s {
		// 下の階
		if string(r) == "U" {
			// 2回経由
			ans += (i - 0) * 2
		} else {
			// 1回経由
			ans += (i - 0)
		}

		// 上の階
		if string(r) == "U" {
			// 1回経由
			ans += (max - i)
		} else {
			// 2回経由
			ans += (max - i) * 2
		}
	}

	fmt.Println(ans)
}
