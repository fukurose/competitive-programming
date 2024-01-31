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
	scanInt()
	s := scanString()
	counter := make(map[string]int, 26)
	for _, c := range strings.Split(s, "") {
		counter[c]++
	}
	ans := 1
	for _, n := range counter {
		ans *= (n + 1) // 選ばないも含める
		ans %= 1000000007
	}
	fmt.Println(ans - 1) // 何も選ばないパターンを除外
}
