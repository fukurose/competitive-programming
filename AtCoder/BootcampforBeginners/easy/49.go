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
	w := scanInt()

	body := make([]string, h)
	for i := range body {
		body[i] = scanString()
	}

	picture := make([][]string, h+2)
	for i := range picture {
		picture[i] = make([]string, w+2)
		for j := range picture[i] {
			picture[i][j] = "#"
		}
	}

	for i := 0; i < h; i++ {
		for j, r := range body[i] {
			picture[i+1][j+1] = string(r)
		}
	}

	for i := range picture {
		for j := range picture[i] {
			fmt.Print(picture[i][j])
		}
		fmt.Println()
	}
}
