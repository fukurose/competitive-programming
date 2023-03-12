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
	n := scanInt()
	k := scanInt()
	x_arr := make([]int, n)
	for i := range x_arr {
		x_arr[i] = scanInt()
	}

	cost := 0
	for _, x := range x_arr {
		if x < k {
			d := min(x-0, k-x)
			cost += (2 * d)
		} else {
			d := x - k
			cost += (2 * d)
		}
	}
	fmt.Println(cost)
}
