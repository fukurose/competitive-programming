package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	n_str := scanString()

	l := len(n_str)
	ans, _ := strconv.Atoi(n_str)

	if l > 1 {
		f, _ := strconv.Atoi(string(n_str[0]))
		r, _ := strconv.Atoi(string(n_str[1:]))
		all9, _ := strconv.Atoi(strings.Repeat("9", l-1))

		ans = 9 * (l - 1)
		if r == all9 {
			ans += f
		} else if f > 1 {
			ans += (f - 1)
		}
	}

	fmt.Println(ans)
}
