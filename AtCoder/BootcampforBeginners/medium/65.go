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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	n := scanInt()
	m := scanInt()
	s := scanString()
	t := scanString()

	l := n * m / gcd(n, m)

	step := l / n * l / m / gcd(l/n, l/m)

	isOK := true
	for i := 0; i/(l/n) < n || i/(l/m) < m; i += step {
		if s[i/(l/n)] != t[i/(l/m)] {
			isOK = false
		}
	}

	if isOK {
		fmt.Println(l)
	} else {
		fmt.Println(-1)
	}
}
