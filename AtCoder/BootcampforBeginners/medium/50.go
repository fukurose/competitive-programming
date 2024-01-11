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
	s := scanInt()
	e := scanInt()

	ans := 2019

	for i := s; i < e; i++ {
		for j := i + 1; j <= e; j++ {
			ans = min(ans, (i%2019)*(j%2019)%2019)
			if ans == 0 {
				break
			}
		}
		if ans == 0 {
			break
		}
	}
	fmt.Println(ans)
}
