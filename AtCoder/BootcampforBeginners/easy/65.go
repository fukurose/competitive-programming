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
	s := scanString()
	length := len(s)

	a_index := 0

	for ; a_index < length; a_index++ {
		if string(s[a_index]) == "A" {
			break
		}
	}

	z_index := length - 1

	for ; z_index > 0; z_index-- {
		if string(s[z_index]) == "Z" {
			break
		}
	}

	fmt.Println(z_index - a_index + 1)
}
