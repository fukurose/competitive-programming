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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	k := scanInt()

	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	dp := make([]int, n-k+1)

	for i := range 

}
