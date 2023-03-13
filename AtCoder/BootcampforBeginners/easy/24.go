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
	a := scanInt()
	b := scanInt()
	m := scanInt()

	a_prices := make([]int, a)
	a_min_price := 1000000
	for i := range a_prices {
		a_prices[i] = scanInt()
		a_min_price = min(a_min_price, a_prices[i])
	}
	b_prices := make([]int, b)
	b_min_price := 1000000
	for i := range b_prices {
		b_prices[i] = scanInt()
		b_min_price = min(b_min_price, b_prices[i])
	}

	min_price := a_min_price + b_min_price

	for i := 0; i < m; i++ {
		a_index := scanInt() - 1
		b_index := scanInt() - 1
		discount := scanInt()

		min_price = min(min_price, a_prices[a_index]+b_prices[b_index]-discount)
	}

	fmt.Println(min_price)
}
