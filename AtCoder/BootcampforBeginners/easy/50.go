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
	a := scanString()
	b := scanString()

	a_len := len(a)
	b_len := len(b)

	ans := "EQUAL"

	if a_len > b_len {
		ans = "GREATER"
	} else if a_len < b_len {
		ans = "LESS"
	} else {
		for i, a_r := range a {
			b_r := b[i]
			n_a, _ := strconv.Atoi(string(a_r))
			n_b, _ := strconv.Atoi(string(b_r))

			if n_a > n_b {
				ans = "GREATER"
				break
			} else if n_a < n_b {
				ans = "LESS"
				break
			}
		}
	}

	fmt.Println(ans)
}
