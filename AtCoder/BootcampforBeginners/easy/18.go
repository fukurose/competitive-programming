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
	s := scanString()

	count := 0
	ans := 0
	for _, r := range s {
		if string(r) == "A" || string(r) == "C" || string(r) == "G" || string(r) == "T" {
			count++
		} else {
			count = 0
		}
		ans = max(ans, count)
	}
	fmt.Println(ans)
}
