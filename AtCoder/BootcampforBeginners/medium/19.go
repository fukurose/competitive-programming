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
	r := scanInt()
	g := scanInt()
	b := scanInt()
	n := scanInt()

	total := 0

	for i := 0; i <= 3000; i++ {
		for j := 0; j <= 3000; j++ {
			sum := (i * r) + (j * g)
			if sum <= n && (n-sum)%b == 0 {
				total++
			}
		}
	}

	fmt.Println(total)

}
