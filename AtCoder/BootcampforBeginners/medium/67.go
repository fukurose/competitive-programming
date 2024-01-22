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

func main() {
	a := scanString()
	b := scanString()
	c := scanString()

	indexA := 0
	indexB := 0
	indexC := 0

	abc := "abc"

	turn := abc[0]
	for {
		if turn == abc[0] {
			if indexA > len(a)-1 {
				break
			}
			turn = a[indexA]
			indexA++
		} else if turn == abc[1] {
			if indexB > len(b)-1 {
				break
			}
			turn = b[indexB]
			indexB++
		} else if turn == abc[2] {
			if indexC > len(c)-1 {
				break
			}
			turn = c[indexC]
			indexC++
		}
	}

	fmt.Println(strings.ToUpper(string(turn)))

}
