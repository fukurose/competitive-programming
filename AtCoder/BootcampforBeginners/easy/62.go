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
	h := scanInt()
	scanInt()

	pixels := make([]string, h)
	for i := range pixels {
		pixels[i] = scanString()
	}

	for i := range pixels {
		fmt.Println(pixels[i])
		fmt.Println(pixels[i])
	}
}
