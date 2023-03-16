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
	w := scanString()

	counter := map[string]int{}

	for _, r := range w {
		counter[string(r)]++
	}

	isBeautiful := true
	for _, v := range counter {
		if v%2 != 0 {
			isBeautiful = false
		}
	}

	if isBeautiful {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
