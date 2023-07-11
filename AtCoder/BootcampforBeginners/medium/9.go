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
	a := scanInt()
	b := scanInt()
	c := scanInt()

	divisor := gcd(a, b)

	if c%divisor == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
