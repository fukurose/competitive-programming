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

type Point struct {
	x int
	y int
}

func main() {
	y := scanInt()
	x := scanInt()

	m := make([][]string, y)
	for i := range m {
		m[i] = make([]string, x)
		m[i] = strings.Split(scanString(), "")
	}

	counter := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] == "#" {
				counter++
			}
		}
	}

	if counter == (x + y - 1) {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}
