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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	n := scanInt()
	k := scanInt()
	r := scanInt()
	s := scanInt()
	p := scanInt()
	t := scanString()

	points := make(map[string]int, 3)
	points["r"] = p
	points["s"] = r
	points["p"] = s

	a := make([]string, n)
	b := strings.Split(t, "")

	ans := 0
	for i := 0; i < k; i++ {
		a[i] = b[i]
		ans += points[b[i]]
	}
	for i := k; i < n; i++ {
		if a[i-k] == b[i] {
			a[i] = "x"
		} else {
			a[i] = b[i]
			ans += points[b[i]]
		}
	}

	fmt.Println(ans)
}
