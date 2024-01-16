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

var max int = 100000000

func main() {
	n := scanInt()
	k := scanInt()
	s := scanInt()

	dummy := (s + 1) % max

	for i := 1; i < n; i++ {
		if k > 0 {
			fmt.Printf("%d ", s)
			k--
		} else {
			fmt.Printf("%d ", dummy)

		}
	}

	if k > 0 {
		fmt.Println(s)
	} else {
		fmt.Println(dummy)
	}
}
