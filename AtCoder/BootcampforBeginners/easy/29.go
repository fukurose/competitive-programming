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

func init() {
	scanner.Split(bufio.ScanWords)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	s := scanString()
	x := 753

	ans := x
	for i := 0; i <= len(s)-3; i++ {
		n, _ := strconv.Atoi(s[i : i+3])
		if n > x {
			ans = min(ans, n-x)
		} else {
			ans = min(ans, x-n)
		}
	}
	fmt.Println(ans)
}
