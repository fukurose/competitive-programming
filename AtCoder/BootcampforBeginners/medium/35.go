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
	if a > b {
		a, b = b, a
	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	a := scanInt()
	b := scanInt()
	c := scanInt()
	d := scanInt()

	cb := b / c
	ca := (a - 1) / c

	c_count := cb - ca

	db := b / d
	da := (a - 1) / d

	d_count := db - da

	l := lcm(c, d)

	lb := b / l
	la := (a - 1) / l
	l_count := lb - la

	ans := b - a + 1

	ans -= c_count
	ans -= d_count
	ans += l_count

	fmt.Println(ans)
}
