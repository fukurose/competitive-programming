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

func isNoBlock(arr []string) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == "#" && arr[i] == "#" {
			return false
		}
	}
	return true
}

func isASkipB(arr []string) bool {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == "." && arr[i] == "." && arr[i+1] == "." {
			return true
		}
	}
	return false
}

func main() {
	scanInt()
	// index で使うので -1 しておく
	a := scanInt() - 1
	b := scanInt() - 1
	c := scanInt() - 1
	d := scanInt() - 1

	s := scanString()
	arr := strings.Split(s, "")

	isBtoD := isNoBlock(arr[b : d+1])
	isAtoC := isNoBlock(arr[a : c+1])

	// A が Bを飛び越す必要がある場合
	if d < c {
		// d の前が空いていないかつ、それまでにスキップ不可の場合
		if arr[d-1] == "#" && !isASkipB(arr[b-1:d+1]) {
			isAtoC = false
		}
	}

	if isBtoD && isAtoC {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
