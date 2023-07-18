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
	n := scanInt()
	k := scanInt()

	var diceP float64 = 1 / float64(n)
	var coinP float64 = 0
	for i := 1; i <= n; i++ {
		count := 0
		for j := i; j < k; j *= 2 {
			count++
		}
		coinP += math.Pow(0.5, float64(count))
	}
	fmt.Println(diceP * coinP)
}
