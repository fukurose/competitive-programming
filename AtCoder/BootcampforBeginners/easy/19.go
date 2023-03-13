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
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	scanInt()
	m := scanInt()
	x := scanInt()

	pass_count := 0
	for i := 0; i < m; i++ {
		n := scanInt()
		if n < x {
			pass_count++
		}
	}

	ans := min(pass_count, m-pass_count)
	fmt.Println(ans)

}
