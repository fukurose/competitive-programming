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
	a := scanInt()
	b := scanInt()

	count := 0
	for i := 1; i < 10; i++ {
		n := (i * 10000) + i
		for j := 0; j < 10; j++ {
			m := n + (j * 1000) + (j * 10)
			for k := 0; k < 10; k++ {
				palindrome := m + (k * 100)
				if a <= palindrome && palindrome <= b {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
