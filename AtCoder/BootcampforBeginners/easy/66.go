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
	table := map[string]bool{"N": false, "W": false, "S": false, "E": false}
	for _, r := range s {
		table[string(r)] = true
	}

	horizontal_check := table["W"] == table["E"]
	vertical_check := table["N"] == table["S"]

	if horizontal_check && vertical_check {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
