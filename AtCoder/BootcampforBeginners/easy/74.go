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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	a := scanInt()
	b := scanInt()
	c := scanInt()

	x := scanInt()
	y := scanInt()

	ans := 0

	// c を2枚買って、a, b を作る方が安い場合
	if (a + b) > (c * 2) {
		// 必要枚数がaの方が多い場合
		if x-y > 0 {
			ans += (c * y * 2)
			// c を2枚買って、a を作る方が安い場合
			if a > c*2 {
				ans += (c * (x - y) * 2)
			} else {
				ans += (a * (x - y))
			}
		} else {
			ans += (c * x * 2)
			// c を2枚買って、a を作る方が安い場合
			if b > c*2 {
				ans += (c * (y - x) * 2)
			} else {
				ans += (b * (y - x))
			}
		}
	} else {
		ans += a * x
		ans += b * y
	}

	fmt.Println(ans)

}
