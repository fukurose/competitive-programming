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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return 0 - a
	}
}

func reverse(a int) int {
	return 0 - a
}

func main() {
	x := scanInt()
	y := scanInt()

	ans := 0

	if abs(x) < abs(y) {
		ans += abs(y) - abs(x)
		if x < 0 {
			// 反転が必要
			ans++
		}

		if y < 0 {
			// 反転が必要
			ans++
		}

	} else if abs(x) > abs(y) {
		ans += abs(x) - abs(y)
		if x > 0 {
			// 反転が必要
			ans++
		}

		if y > 0 {
			// 反転が必要
			ans++
		}

	} else {
		// 反転が必要
		ans++
	}
	fmt.Println(ans)
}
