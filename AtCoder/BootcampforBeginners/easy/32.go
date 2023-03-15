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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	scanString()
	s := scanString()

	x := 0
	ans := x
	for _, r := range s {
		switch string(r) {
		case "I":
			x++
		case "D":
			x--
		}
		ans = max(ans, x)
	}

	fmt.Println(ans)

}
