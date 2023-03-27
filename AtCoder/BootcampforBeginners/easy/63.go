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
	bricks := make([]int, n)

	for i := range bricks {
		bricks[i] = scanInt()
	}

	next_num := 1
	broken_counter := 0

	for _, brick := range bricks {
		if brick == next_num {
			next_num++
		} else {
			broken_counter++
		}
	}
	// 全部破壊した場合
	if broken_counter == n {
		broken_counter = -1
	}

	fmt.Println(broken_counter)

}
