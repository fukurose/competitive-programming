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
	n := scanInt()
	a := scanInt()
	b := scanInt()

	// ab間にあるマス目
	distBetweenAB := (b - a - 1)

	ans := 0
	if distBetweenAB%2 == 0 {
		ans += min(a-1, n-b) // 近い方の端に移動
		// 端で動かず、試合をしてAB間にあるマス目を奇数にする
		if distBetweenAB > 0 {
			ans++
			distBetweenAB--
		}
	}
	ans += distBetweenAB/2 + 1
	fmt.Println(ans)
}
