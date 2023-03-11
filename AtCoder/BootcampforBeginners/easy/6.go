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
	h := scanInt()
	w := scanInt()

	squares := h * w

	if h == 1 || w == 1 {
		fmt.Println(1)
	} else if squares%2 == 0 {
		fmt.Println(squares / 2)
	} else {
		fmt.Println((squares / 2) + 1)
	}
}
