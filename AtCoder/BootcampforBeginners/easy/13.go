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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	a := scanInt()
	b := scanInt()
	c := scanInt()

	count := 0

	for a%2 == 0 && b%2 == 0 && c%2 == 0 {
		if a == b && b == c {
			count = -1
			break
		}
		next_a := b/2 + c/2
		next_b := a/2 + c/2
		next_c := a/2 + b/2

		count++

		a = next_a
		b = next_b
		c = next_c

	}

	fmt.Println(count)
}
