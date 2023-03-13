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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()

	count := 0
	for i := 2; i <= n; i *= 2 {
		count++
	}

	if count == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(math.Pow(2, float64(count)))
	}

}
