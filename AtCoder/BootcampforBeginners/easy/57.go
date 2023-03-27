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
	price := 1.0
	ans := -1

	for {
		tax8 := int(price * 0.08)
		tax10 := int(price * 0.1)
		if tax8 == a && tax10 == b {
			ans = int(price)
			break
		} else if tax8 > a || tax10 > b {
			break
		}
		price += 1.0
	}

	fmt.Println(ans)
}
