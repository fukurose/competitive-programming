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
	k := scanInt()
	s := scanInt()

	ans := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - (x + y)
			if z >= 0 && z <= k {
				ans++

			}
		}
	}

	fmt.Println(ans)
}
