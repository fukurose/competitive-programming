package main

import (
	"bufio"
	"fmt"
	"math"
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
	scanInt()
	k := scanInt()

	ans := a - b

	if k%2 != 0 {
		ans *= -1
	}

	if math.Abs(float64(ans)) > math.Pow(10, 18) {
		fmt.Println("Unfair")
	} else {
		fmt.Println(ans)
	}
}
