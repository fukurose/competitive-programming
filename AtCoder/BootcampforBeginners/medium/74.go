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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	n := scanInt()
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, 26) // a ~ z
		s := scanString()
		for _, r := range s {
			arr[i][r-97]++
		}
	}

	// a ~ z
	for i := 0; i < 26; i++ {
		m := arr[0][i]
		for j := 1; j < n; j++ {
			m = min(m, arr[j][i])
		}
		s := string(rune(i + 97))
		for k := 0; k < m; k++ {
			fmt.Print(s)
		}
	}
	fmt.Println()
}
