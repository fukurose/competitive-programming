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
	k := scanInt()

	indexOfNot1 := 0
	for i, r := range s {
		if string(r) != "1" {
			indexOfNot1 = i
			break
		}
	}

	if k > (indexOfNot1) {
		fmt.Println(string(s[indexOfNot1]))
	} else {
		fmt.Println(1)
	}
}
