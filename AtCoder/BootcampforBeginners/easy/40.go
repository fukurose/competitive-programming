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

func main() {
	a := scanInt()
	b := scanInt()

	ans := "Zero"
	if a > 0 && b > 0 {
		ans = "Positive"
	} else if a < 0 && b < 0 {
		diff := (b * -1) - (a * -1) + 1
		if diff%2 == 0 {
			ans = "Positive"
		} else {
			ans = "Negative"
		}
	}

	fmt.Println(ans)

}
