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

func main() {
	x := scanInt()

	prices := []int{50, 10, 5, 1}

	limits := make([]int, 4)
	for i := range limits {
		limits[i] = scanInt()
	}

	ans := 0
	for i, price := range prices {
		if x < price {
			continue
		}

		count := limits[i]
		if x/price < limits[i] {
			count = x / price
		}
		x -= price * count
		ans += count
	}

	fmt.Println(ans)
}
