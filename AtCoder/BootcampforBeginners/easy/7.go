package main

import (
	"bufio"
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
	bingo := make([][]int, 3)
	for i := range bingo {
		bingo[i] = make([]int, 3)
	}

}
